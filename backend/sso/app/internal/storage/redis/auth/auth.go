package auth

import (
	"context"
	"github.com/redis/go-redis/v9"
	"vss/sso/internal/storage"

	"github.com/pkg/errors"
	"vss/sso/internal/config"
	rdconnector "vss/sso/pkg/connectors/redis_connector"
	logger "vss/sso/pkg/logger/handlers/slogpretty"
)

type RedisCache struct {
	client *redis.Client
}

func New(ctx context.Context) (*RedisCache, error) {
	client, err := rdconnector.Connect(ctx, config.Get().Redis, logger.Log)
	if err != nil {
		return nil, errors.Wrapf(err,
			"Storage.Redis.Auth.New %s",
			"cannot create connection to redis")
	}

	return &RedisCache{
		client: client,
	}, nil
}

func (r *RedisCache) UpdateSessionToken(ctx context.Context, args storage.SessionToken) error {
	_, err := r.client.Set(ctx, args.UserID, args.AccessToken, args.TTL).Result()
	if err != nil {
		return errors.Wrapf(err, "could not update session token")
	}

	return nil
}
