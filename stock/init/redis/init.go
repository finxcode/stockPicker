package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"stockPicker/stock/init/config"
)

func InitRedis(config *config.Config) (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
		PoolSize: config.Redis.PoolSize,
	})
	ctx := context.Background()
	if _, err := rds.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	return rds, nil
}
