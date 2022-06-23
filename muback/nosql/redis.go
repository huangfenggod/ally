package nosql

import (
	"context"
	"github.com/go-redis/redis/v8"
	"muback/config"
	"time"
)

var Client *redis.Client
func InitRedis()  {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.RedisNetwork,
		Password: config.Cfg.Redis.RedisPassword,
		DB:       0})
	Client = client
}

func SetTokenInRedis(token string)bool  {
	ctx := context.Background()
	err := Client.Set(ctx, token, nil, time.Hour*2).Err()
	if err!=nil {
		return false
	}
	return true
}

func ExistsToken(token string)bool{
	exists := Client.Exists(context.Background(), token)
	if exists.Val()>0 {
		return true
	}
	return false
}

func RemoveToken(token string) bool{
	del := Client.Del(context.Background(),token)
	if del.Err()!=nil {
		return false
	}
	return true
}
