package _select

import (
	"context"
	"fmt"
	"time"
)

func SelectUse1() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	quiteCh := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- struct{}{}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- struct{}{}
	}()

	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println("ch1")
				quiteCh <- struct{}{}
			case <-ch2:
				fmt.Println("ch2")
				quiteCh <- struct{}{}
			default:
				fmt.Println("default")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	<-quiteCh
	fmt.Println("complete")
}

func SelectForTimeout() {
	ctx, cancleFn := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancleFn()

	ch1 := make(chan struct{})
	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("timeout: %+v\n", ctx.Err())
	case <-ch1:
		fmt.Println("ch1 done")
	}

	fmt.Println("complete")
}

func doSomething(ctx context.Context, ts time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("do something ctx done:", ctx.Err())
	case <-time.After(ts):
		fmt.Println("do something complete")
	}
}

func SelectForTimeout2() {
	ctx, cancle := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancle()

	doSomething(ctx, 1500*time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("select timeout2:", ctx.Err())
	}
}
