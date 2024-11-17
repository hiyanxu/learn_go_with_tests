package redis

import (
	//"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

/**
测试go-redis包.
 */

func TryCommand2() {
	client := redis.NewClient(&redis.Options{
		Addr:               "localhost:6379",
	})
	//ctx := context.Background()

	isExist, err := client.Exists("h1").Result()
	fmt.Println(isExist, err)

	// 输出：
	// true <nil>
	// false <nil>
	resp, err := client.SetNX("k1", "v1", 60*time.Second).Result()
	fmt.Println(resp, err)

	resp2, err := client.SetNX("k1", "v2", 60*time.Second).Result()
	fmt.Println(resp2, err)
}
