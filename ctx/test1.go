package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "k1", "v1")
	wg := sync.WaitGroup{}
	wg.Add(2)
	fmt.Printf("地址：%p\n", &ctx)
	fmt.Println(ctx.Value("k1"))

	go func(ctx2 context.Context) {
		fmt.Printf("ctx2: %p, k1: %s\n", &ctx, ctx.Value("k1"))
		wg.Done()
	}(ctx)

	go func(ctx context.Context) {
		fmt.Printf("ctx3: %p, k1: %s\n", &ctx, ctx.Value("k1"))
		wg.Done()
	}(ctx)

	wg.Wait()
}
