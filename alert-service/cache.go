package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cacher interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
}

type Redis struct {
	client *redis.Client
}

func NewRedis(addr string) (Cacher, error) {
	opt, err := redis.ParseURL(addr)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Set(ctx context.Context, key string, value string) error {
	return r.client.Set(ctx, key, value, 0).Err()
}
