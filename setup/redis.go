package setup

import (
	"github.com/redis/go-redis/v9"
)

var REDIS *redis.Client

func RunRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     VarArgData.RedisHost,
		Password: VarArgData.RedisPass,
		DB:       VarArgData.RedisDB,
	})

	REDIS = rdb
}
