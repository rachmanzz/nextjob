package setup

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var REDIS *redis.Client

var ctx = context.Background()

func RunRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     VarArgData.RedisHost,
		Password: VarArgData.RedisPass,
		DB:       VarArgData.RedisDB,
	})

	startAt := time.Now()
	err := rdb.Set(ctx, "nextjob_up", startAt.String(), 0).Err()
	if err != nil {
		panic("not connect to redis server")
	}

	REDIS = rdb
}
