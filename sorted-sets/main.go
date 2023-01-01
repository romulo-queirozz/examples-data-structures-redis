package main

import (
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v9"
	"math/rand"
	"redis"
	"strconv"
)

var ctx = context.Background()

func main() {
	key := "players"
	client := redis.MyRedisClient()

	for i := 0; i < len(users()); i++ {
		player := users()[i]
		client.ZAdd(ctx, key, redis2.Z{Member: player, Score: float64(rand.Intn(1000))})
	}

	players := client.ZRevRangeWithScores(ctx, key, 0, int64(len(users()))).Val()

	for i, player := range players {
		player := player.Member
		score := players[i].Score
		fmt.Println(strconv.Itoa(i+1), "Place:", player, "| Score", score)
	}
}

func users() map[int]string {
	return map[int]string{
		0: "Rob",
		1: "Lilly",
		2: "Bob",
		3: "Anna",
		4: "Antony",
	}
}
