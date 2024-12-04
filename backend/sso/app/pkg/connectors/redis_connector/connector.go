package redis_connector

import (
	"context"
	"github.com/pkg/errors"

	Client "github.com/redis/go-redis/v9"
)

type Log interface {
	Info(string)
	Error(string)
}

func Connect(ctx context.Context, cfg Redis, log Log) (*Client.Client, error) {
	conn, err := NewRedisClient(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		log.Info("Closing redis connection")
		if err := conn.Close(); err != nil {
			log.Error(err.Error())
		} else {
			log.Info("Redis connection closed properly")
		}
	}()

	return conn, nil
}

func NewRedisClient(ctx context.Context, cfg Redis) (*Client.Client, error) {
	opts, err := cfg.BuildOptions()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to build connection options for redis connector")
	}

	client := Client.NewClient(opts)
	result := client.Ping(ctx)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return client, nil
}
