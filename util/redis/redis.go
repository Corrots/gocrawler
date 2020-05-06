package redis

import (
	"github.com/go-redis/redis/v7"
	"log"
	"time"
)

const expiration  = time.Duration(86400)

var (
	client *redis.Client
)

func init()  {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     "0.0.0.0:6379",
			Password: "",
			DB:       0,
		})
		_, err := client.Ping().Result()
		if err != nil {
			log.Printf("redis init err: %v\n", err)
			panic(err)
		}
	}
}

func Set(key string, value interface{}) error {
	return client.Set(key, value, 0).Err()
}

func Get(key string) (string, error) {
	return client.Get(key).Result()
}

func KeyExists(key string) bool {
	_, err := client.Exists(key).Result()
	if err != nil || err == redis.Nil {
		return false
	}
	return true
}

func SAdd(key string, members string) error {
	return client.SAdd(key, members).Err()
}

func SIsMember(key , value string) bool {
	return client.SIsMember(key, value).Val()
}

func SMember(key string) (interface{}, error) {
	return client.SMembers(key).Result()
}

func LPush(key string, values ...interface{}) error {
	return client.LPush(key, values).Err()
}

func LPop(key string) (interface{}, error) {
	return client.LPop(key).Result()
}

func LLen(key string) int64 {
	return client.LLen(key).Val()
}