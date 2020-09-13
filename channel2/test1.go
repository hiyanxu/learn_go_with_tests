package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg = sync.WaitGroup{}

type a struct {
	ready chan bool
}

func (a *a) work() error {
	time.Sleep(3 * time.Second)
	return errors.New("err")
	//return nil
}

func (a *a) setup() error {
	close(a.ready)
	return nil
}

func Start() error {
	a1 := &a{ready: make(chan bool)}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("开始执行defer...")

		for {
			a1.setup()
			if err := a1.work(); err != nil {
				fmt.Printf("出现error: %s\n", err)
			}

			fmt.Printf("dddddd\n")

			// 若worker失败，第一次执行到该处，下面a1.ready会取出channel中的数据，会出现"start and running..."；
			// 之后for循环继续执行，若再出现失败，走到该channel处写入，由于没有接收的地方，会导致一直阻塞在该处
			//a1.ready <- true
			a1.ready = make(chan bool)
		}
	}()

	<-a1.ready
	fmt.Printf("start and running...\n")

	ctx, cancel := context.WithCancel(context.Background())
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()

	return nil
}

//func main() {
//	Start()
//}
