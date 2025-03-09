package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Sesuaikan dengan Redis server-mu
	// DB:   0,                // Gunakan database default (0)
})

func RateLimiting(user string) bool {
	key := "rate :" + user

	count, _ := rdb.Incr(ctx, key).Result()
	rdb.Expire(ctx, key, time.Minute)
	return count <= 10 //max 10 request data
}
