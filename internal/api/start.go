package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/1nterdigital/aka-im-discover/internal/api/mw"
	"github.com/1nterdigital/aka-im-discover/internal/api/util"
	"github.com/1nterdigital/aka-im-discover/internal/repository"
	"github.com/1nterdigital/aka-im-discover/internal/service"
	"github.com/1nterdigital/aka-im-discover/internal/usecase"
	"github.com/1nterdigital/aka-im-discover/pkg/common/config"
	"github.com/1nterdigital/aka-im-discover/pkg/common/db"
	"github.com/1nterdigital/aka-im-discover/pkg/common/db/database"
	"github.com/1nterdigital/aka-im-discover/pkg/common/imapi"
	"github.com/1nterdigital/aka-im-discover/pkg/common/kdisc"
	disetcd "github.com/1nterdigital/aka-im-discover/pkg/common/kdisc/etcd"
	"github.com/1nterdigital/aka-im-discover/pkg/common/tokenverify"
	"github.com/1nterdigital/aka-im-tools/db/mysqlutil"
	"github.com/1nterdigital/aka-im-tools/db/redisutil"
	"github.com/1nterdigital/aka-im-tools/discovery/etcd"
	"github.com/1nterdigital/aka-im-tools/errs"
	"github.com/1nterdigital/aka-im-tools/log"
	"github.com/1nterdigital/aka-im-tools/system/program"
	"github.com/1nterdigital/aka-im-tools/utils/datautil"
	"github.com/1nterdigital/aka-im-tools/utils/runtimeenv"
)

type Config struct {
	ApiConfig      config.API
	Discovery      config.Discovery
	Share          config.Share
	Admin          config.Admin
	RedisConfig    config.Redis
	PostgresConfig config.Postgres
	MysqlConfig    config.Mysql
	TracerConfig   config.Tracer
	RuntimeEnv     string
}

type discoverService struct {
	Database database.DiscoverDatabaseInterface
	Token    *tokenverify.Token
}

// initDatabase creates gorm + mysql client
func initDatabase(ctx context.Context, cfg *Config) (*gorm.DB, *mysqlutil.Client, error) {
	pgDB, err := mysqlutil.NewMysqlDB(ctx, cfg.MysqlConfig.Build())
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: pgDB.DB}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	if err = db.InitiateTable(gormDB); err != nil {
		return nil, nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	return gormDB, pgDB, nil
}

// initService wires up repository, usecase, and discover service
func initService(
	cfg *Config, conn *gorm.DB, pgDB *mysqlutil.Client, rdb redis.UniversalClient,
) (*discoverService, *usecase.UseCase, error) {
	repo := repository.NewRepository(conn)
	uc, err := usecase.New(repo)
	if err != nil {
		return nil, nil, err
	}

	srv := &discoverService{
		Token: &tokenverify.Token{
			Expires: time.Duration(cfg.Admin.TokenPolicy.Expire) * 24 * time.Hour,
			Secret:  cfg.Admin.Secret,
		},
	}

	srv.Database, err = database.NewDiscoverDatabase(pgDB, rdb, srv.Token)
	if err != nil {
		return nil, nil, err
	}

	return srv, uc, nil
}

// setupServer configures Gin + HTTP server
func setupServer(cfg *Config, discoverApi *service.Api, mwApi *mw.MW, apiPort int) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	engine := SetRouter(cfg.TracerConfig.AppName.Api, discoverApi, mwApi)

	return &http.Server{
		Addr:              fmt.Sprintf(":%d", apiPort),
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func Start(ctx context.Context, index int, cfg *Config) (err error) {
	log.CInfo(ctx, "Starting DISCOVER-API server3")
	cfg.RuntimeEnv = runtimeenv.PrintRuntimeEnvironment()
	if len(cfg.Share.DiscoverAdmin) == 0 {
		return errs.New("share discover admin not configured")
	}

	// Redis
	rdb, err := redisutil.NewRedisClient(ctx, cfg.RedisConfig.Build())
	if err != nil {
		return err
	}

	// DB
	conn, pgDB, err := initDatabase(ctx, cfg)
	if err != nil {
		return err
	}

	// Service + usecase
	srv, uc, err := initService(cfg, conn, pgDB, rdb)
	if err != nil {
		return err
	}

	// Discovery client
	client, err := kdisc.NewDiscoveryRegister(&cfg.Discovery, cfg.RuntimeEnv, nil)
	if err != nil {
		return err
	}

	// External dependencies
	im := imapi.New(cfg.Share.AkaIM.ApiURL, cfg.Share.AkaIM.Secret, cfg.Share.AkaIM.AdminUserID)
	base := util.Api{
		ImUserID:            cfg.Share.AkaIM.AdminUserID,
		ProxyHeader:         cfg.Share.ProxyHeader,
		DiscoverAdminUserID: cfg.Share.DiscoverAdmin[0],
	}

	// API + middleware
	discoverApi := service.New(cfg.TracerConfig.AppName.Api, im, &base, *uc)
	mwApi := mw.New(srv.Token, srv.Database)

	// HTTP server
	apiPort, err := datautil.GetElemByIndex(cfg.ApiConfig.Api.Ports, index)
	if err != nil {
		return err
	}
	server := setupServer(cfg, discoverApi, mwApi, apiPort)

	// Run server
	netDone := make(chan struct{}, 1)
	var netErr error
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			netErr = errs.WrapMsg(err, fmt.Sprintf("api start err: %s", server.Addr))
			netDone <- struct{}{}
		}
	}()

	// Config watcher
	if cfg.Discovery.Enable == kdisc.ETCDCONST {
		cm := disetcd.NewConfigManager(client.(*etcd.SvcDiscoveryRegistryImpl).GetClient(),
			[]string{
				config.DiscoverApiCfgFileName,
				config.DiscoveryConfigFileName,
				config.ShareFileName,
				config.LogConfigFileName,
			},
		)
		cm.Watch(ctx)
	}

	// Graceful shutdown
	timeoutShutdown := 15 * time.Second
	return gracefulShutdown(server, timeoutShutdown, netDone, netErr)
}

func shutdown(server *http.Server, timeout time.Duration) func() error {
	return func() error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			return errs.WrapMsg(err, "shutdown err")
		}
		return nil
	}
}

func gracefulShutdown(server *http.Server, timeout time.Duration, netDone chan struct{}, netErr error) error {
	sd := shutdown(server, timeout)
	disetcd.RegisterShutDown(sd)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(sigs)

	select {
	case <-sigs:
		log.CInfo(context.Background(), "received shutdown signal, stopping server...")
		program.SIGTERMExit()
		return sd()
	case <-netDone:
		close(netDone)
		return netErr
	}
}
