package main

import (
	"context"
	"fmt"
	"redis"
)

var ctx = context.Background()

func main() {
	client := redis.MyRedisClient()

	key := "count-page-home-view"

	client.PFAdd(ctx, key, "user-1", "user-1")
	client.PFAdd(ctx, key, "user-1", "user-2", "user-3")
	client.PFAdd(ctx, key, "user-4", "user-10")
	client.PFAdd(ctx, key, "user-6", "user-1", "user-1")

	count := client.PFCount(ctx, key).Val()
	fmt.Println("Page Views:", count)
}
