package database

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/1nterdigital/aka-im-discover/pkg/common/db/cache"
	"github.com/1nterdigital/aka-im-discover/pkg/common/tokenverify"
	"github.com/1nterdigital/aka-im-tools/db/mysqlutil"
)

type (
	DiscoverDatabaseInterface interface {
		GetTokens(ctx context.Context, userID string) (map[string]int32, error)
	}

	DiscoverDatabase struct {
		cache   cache.TokenInterface
		mysqlDB *mysqlutil.Client
	}
)

func NewDiscoverDatabase(
	mysqlCli *mysqlutil.Client,
	rdb redis.UniversalClient,
	token *tokenverify.Token,
) (DiscoverDatabaseInterface, error) {
	return &DiscoverDatabase{
		mysqlDB: mysqlCli,
		cache:   cache.NewTokenInterface(rdb, token),
	}, nil
}

func (o *DiscoverDatabase) GetTokens(ctx context.Context, userID string) (map[string]int32, error) {
	return o.cache.GetTokensWithoutError(ctx, userID)
}
