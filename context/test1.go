package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 返回一个new(emptyCtx).
	ctx := context.Background()
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("ctx done:", ctx.Err())
		}
	}()

	ctx2, _ := context.WithCancel(ctx)
	go func() {
		select {
		case <-ctx2.Done():
			fmt.Println("ctx2 done:", ctx2.Err())
		}
	}()
	ctx3, cancel3 := context.WithCancel(ctx2) // 返回的 cancel3 函数，用于取消 ctx3 以及从它衍生出来的子上下文.
	go func() {
		select {
		case <-ctx3.Done():
			fmt.Println("ctx3 done:", ctx3.Err())
		}
	}()
	cancel3()
	time.Sleep(1 * time.Second)
}
