package main

import "context"

func main() {
	// 返回一个new(emptyCtx).
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	ctx, cancel = context.WithCancel(ctx)
	cancel()

}
