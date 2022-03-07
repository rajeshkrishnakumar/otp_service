package redis

import (
	"context"
	"fmt"
	"time"

	"otp_service/utils"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client
var ctx = context.Background()

func Connect() {
	client = redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    utils.ApplicationConfig.REDIS_HOST + ":" + utils.ApplicationConfig.REDIS_PORT,
		DB:      utils.ApplicationConfig.REDIS_DB,
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)
	fmt.Println("Redis connected")
}

func Getdata(key string) (ok bool, response string) {
	resultFromRedis, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, "no_data"
	}
	if err != nil {
		return false, err.Error()
	}
	return true, resultFromRedis
}

func SetDataWithTTL(key string, value interface{}, ttl time.Duration) bool {
	err := client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func SetData(key string, value interface{}) (ok bool, response string) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false, err.Error()
	}
	return true, "data saved"
}

func DelData(key string) (ok bool, response string) {
	err := client.Del(ctx, key).Err()
	if err != nil {
		fmt.Println(err)
		return false, err.Error()
	}
	return true, "data deleted"
}

func KeysData(pattern string) []string {
	return client.Keys(ctx, pattern).Val()
}

func GetExpire(key string) time.Duration {
	return client.TTL(ctx, key).Val()
}

func SetExpire(key string, ttl time.Duration) bool {
	return client.Expire(ctx, key, ttl).Val()
}
