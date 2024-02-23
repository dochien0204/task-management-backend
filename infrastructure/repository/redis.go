package repository

import (
	"context"
	"fmt"
	"source-base-go/config"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis(c config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", c.Host, c.Port),
	})

	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
