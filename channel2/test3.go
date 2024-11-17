package main

import "fmt"

/*
用 golang channel 实现生产者和消费者模型.
*/
func Handle() {
	ch1 := make(chan int)
	done := make(chan struct{})

	// 生产者.
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}

		fmt.Println("开始 close ch1")
		close(ch1)
		fmt.Println("close ch1")
	}()

	// 消费者.
	go func() {
		//for {
		//	select {
		//	case val, ok := <-ch1:
		//		if ok {
		//			fmt.Printf("consumer get val: %d\n", val)
		//		} else {
		//			fmt.Println("消费完毕")
		//			close(done)
		//			return
		//		}
		//	default:
		//
		//	}
		//}

		// 采用 for range 方式.
		for val := range ch1 {
			fmt.Printf("consumer get by range: %d\n", val)
		}
		close(done)
	}()

	<-done
	fmt.Println("done")
	<-done
	fmt.Println("又读了一次 done")
}

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		tmpi := i
		go func() {
			ch <- tmpi
		}()
	}
}

func consumer(ch chan int, done chan struct{}) {
	completes := make(chan struct{}, 10)
	for i := 1; i <= 10; i++ {
		// 启动 10 个消费者 goroutine 去消费.
		go func() {
			select {
			case val := <-ch:
				fmt.Printf("消费者消费到: %d\n", val)
				completes <- struct{}{}
			}
		}()
	}

	for range completes {
		
	}
	close(done)
}

// Handle2 带缓冲的 channel 实现生产者消费者模型.
func Handle2() {
	ch1 := make(chan int, 10)
	done := make(chan struct{})
	go producer(ch1)
	go consumer(ch1, done)

	<-done
	fmt.Printf("消费完成\n")
}
