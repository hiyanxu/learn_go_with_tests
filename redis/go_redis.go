package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

/**
测试go-redis包.
 */

func TryCommand2() {
	client := redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
	})
	ctx := context.Background()

	isExist, err := client.Exists(ctx, "h1").Result()
	fmt.Println(isExist, err)
}
