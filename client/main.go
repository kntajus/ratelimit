package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	_ = rdb.FlushDB(ctx).Err()

	limiter := redis_rate.NewLimiter(rdb)

	ticker := time.NewTicker(250 * time.Millisecond)
	defer ticker.Stop()

	id := os.Getenv("CLIENT_ID")
	client := http.Client{}
	for range ticker.C {
		res, err := limiter.Allow(ctx, "apicall", redis_rate.PerSecond(6))
		if err != nil {
			panic(err)
		}
		if res.Allowed > 0 {
			client.Get("http://server:8080/?id=" + id)
		} else {
			fmt.Println("Rejected")
		}
	}
}
