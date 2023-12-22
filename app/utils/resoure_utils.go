package utils

import (
	"context"
	"github.com/edwardzhanged/novel-go/app/conf"
	"time"
)

type RedisVerifyStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

var Store = &RedisVerifyStore{
	Expiration: time.Second * 180,
	PreKey:     "CAPTCHA_",
	Context:    context.TODO(),
}

func (r *RedisVerifyStore) Set(id string, value string) error {
	conf.GbRedis.Set(r.Context, r.PreKey+id, value, r.Expiration)
	return nil
}

func (r *RedisVerifyStore) Get(id string, clear bool) string {
	val, err := conf.GbRedis.Get(r.Context, r.PreKey+id).Result()
	if clear {
		conf.GbRedis.Del(r.Context, id)
	}
	if err != nil {
		return ""
	}
	return val
}

func (r *RedisVerifyStore) Verify(id string, answer string, clear bool) bool {
	val, err := conf.GbRedis.Get(r.Context, r.PreKey+id).Result()
	if err != nil {
		return false
	}
	conf.GbRedis.Del(r.Context, r.PreKey+id)
	return val != "" && val == answer
}
