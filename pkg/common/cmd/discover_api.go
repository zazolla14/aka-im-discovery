package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/1nterdigital/aka-im-discover/internal/api"
	"github.com/1nterdigital/aka-im-discover/pkg/common/config"
	"github.com/1nterdigital/aka-im-discover/pkg/common/constant"
	"github.com/1nterdigital/aka-im-tools/log"
	"github.com/1nterdigital/aka-im-tools/system/program"
	"github.com/1nterdigital/aka-im-tools/tracer"
)

type DiscoverApiCmd struct {
	*RootCmd
	ctx       context.Context
	configMap map[string]any
	apiConfig api.Config
}

func NewDiscoverApiCmd() *DiscoverApiCmd {
	var ret DiscoverApiCmd
	ret.configMap = map[string]any{
		config.ShareFileName:           &ret.apiConfig.Share,
		config.DiscoverApiCfgFileName:  &ret.apiConfig.ApiConfig,
		config.DiscoveryConfigFileName: &ret.apiConfig.Discovery,
		config.AdminFileName:           &ret.apiConfig.Admin,
		config.RedisConfigFileName:     &ret.apiConfig.RedisConfig,
		config.MysqlConfigFileName:     &ret.apiConfig.MysqlConfig,
		config.TracerConfigFileName:    &ret.apiConfig.TracerConfig,
	}
	ret.RootCmd = NewRootCmd(program.GetProcessName(), WithConfigMap(ret.configMap))
	ret.ctx = context.WithValue(context.Background(), constant.ContextKeyVersion, config.Version)
	ret.Command.RunE = func(_ *cobra.Command, _ []string) error {
		return ret.runE()
	}

	return &ret
}

func (a *DiscoverApiCmd) Exec() error {
	return a.Execute()
}

func (a *DiscoverApiCmd) runE() error {
	if a.apiConfig.TracerConfig.Enable {
		shutdownTracer, err := tracer.InitTracer(a.ctx,
			a.apiConfig.TracerConfig.AppName.Api,
			a.apiConfig.TracerConfig.Otel.Collector.Address,
		)
		if err != nil {
			return err
		}

		defer func() {
			if errx := shutdownTracer(a.ctx); errx != nil {
				log.ZError(a.ctx, "an error occurred while shutdown", errx)
			}
		}()
	}

	return api.Start(a.ctx, a.Index(), &a.apiConfig)
}
