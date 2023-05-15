package setup

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var REDIS *redis.Client

func RunRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"), // "localhost:6379"
		Password: "",                      // no password set
		DB:       0,                       // use default DB
	})

	REDIS = rdb
}
