package cache

import (
	"context"
	"fmt"
	"social-network-4/config"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Config("CACHE_ADDR"),
		Password: config.Config("CACHE_PASS"),
	})
	if rdb == nil {
		panic("Failed to start Cache")
	}
	fmt.Println("Connected to Cache!")
}
