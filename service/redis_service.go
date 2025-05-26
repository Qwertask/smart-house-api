package service

import (
	"SmartHouseAPI/config"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisService(ctx context.Context, config config.RedisConfig) (*RedisService, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr: config.Address,
		DB:   config.Database,
	})

	err := client.Ping(timeoutCtx).Err()
	if err != nil {
		log.Printf("Redis connection failed: %v", err)
		return nil, err
	}

	return &RedisService{client: client}, nil
}

func (c RedisService) Set(ctx context.Context, key string) error {
	err := c.client.Set(ctx, key, "", time.Hour*24*30).Err()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("redis set failed: %v", err)
	}
	return nil
}

func (c RedisService) Get(ctx context.Context, key string) (bool, error) {
	_, err := c.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		log.Println(err)
		return false, nil
	}
	if err != nil {
		log.Println(err)
		return false, fmt.Errorf("redis get failed: %v", err)
	}

	return true, nil
}
