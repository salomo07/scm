package services

import (
	"context"
	"scm/config"

	"github.com/redis/go-redis/v9"
)

func SaveValueRedis(key string, value string) {
	var ctx = context.Background()
	opt, _ := redis.ParseURL(config.GetCredRedis())
	client := redis.NewClient(opt)

	client.Set(ctx, key, value, 0)
	print(key + " is saved")
}

func GetValueRedis(key string) string {
	var ctx = context.Background()

	opt, _ := redis.ParseURL(config.GetCredRedis())
	client := redis.NewClient(opt)
	var res = client.Get(ctx, key).Val()
	return res
}
