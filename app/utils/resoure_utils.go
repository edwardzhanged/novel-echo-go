package utils

import (
	"context"
	"fmt"
	"github.com/edwardzhanged/novel-go/app/conf"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisVerifyStore struct {
	client *redis.Client
}

func (r *RedisVerifyStore) Set(id string, value string) error {
	r.client.Set(context.TODO(), id, value, 120*time.Second)
	return nil
}

func (r *RedisVerifyStore) Get(id string, clear bool) string {
	val, err := r.client.Get(context.TODO(), id).Result()
	if clear {
		r.client.Del(context.TODO(), id)
	}
	if err != nil {
		return ""
	}
	return val
}

func (r *RedisVerifyStore) Verify(id string, answer string, clear bool) bool {
	val, err := r.client.Get(context.TODO(), id).Result()
	if err != nil {
		return false
	}
	return val != "" && val == answer
}

func NewRedisVerifyStore() *RedisVerifyStore {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.GbViper.GetString("redis.host"),
			conf.GbViper.GetString("redis.port")),
		Password: conf.GbViper.GetString("redis.password"),
	})

	return &RedisVerifyStore{client: client}
}

var Store = NewRedisVerifyStore()
