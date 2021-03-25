package main

import (
	"context"
	"fmt"
	"time"
)

//func t1(c1 chan int, c2 chan int) {
//	for {
//		select {
//		case <-c1:
//			fmt.Println("c1 out")
//		case <-c2:
//			fmt.Println("c2 out")
//		case <-time.After(time.Second * 1):
//			fmt.Println("timeout")
//		}
//	}
//
//	fmt.Println("select complete")
//}
//
//func main() {
//	c1 := make(chan int)
//	c2 := make(chan int)
//	go func() {
//		t1(c1, c2)
//	}()
//	fmt.Println("select调用 阻塞")
//	time.Sleep(2 * time.Second)
//	go func() {
//		c1 <- 1
//	}()
//	time.Sleep(3 * time.Second)
//}

// 采用chan通知结束
//func main() {
//	stop := make(chan bool)
//	go func() {
//		for {
//			select {
//			case <-stop:
//				fmt.Println("监控结束")
//				return
//			default:
//				fmt.Println("监控中。。。。。")
//				time.Sleep(2 * time.Second)
//			}
//		}
//	}()
//
//	time.Sleep(10 * time.Second)
//	fmt.Println("通知去结束")
//	stop <- true
//	time.Sleep(5 * time.Second)
//}

/**
测试通过chan通知   测试直接赋值一个已经关闭的chan，看select是否可以接收到通知.
以下程序输出：
waiting...
waiting...
waiting...
waiting...
waiting...
waiting...
close..

*/
//var closeChan = make(chan struct{})
//
//func init() {
//	close(closeChan)
//
//}
//
//type A struct {
//	c1 chan struct{}
//}
//
//func main() {
//	a := new(A)
//	go func() {
//		for {
//			select {
//			case <-a.c1:
//				fmt.Println("close..")
//				return
//			default:
//				fmt.Println("waiting...")
//				time.Sleep(1 * time.Second)
//			}
//		}
//
//	}()
//
//	time.Sleep(5 * time.Second)
//	if a.c1 == nil {
//		a.c1 = closeChan // 直接赋值一个已经被关闭的chan，在上面select处，是可以接收到一个关闭的消息的
//	}
//	time.Sleep(3 * time.Second)
//
//}

// 采用ctx通知结束
func main() {
	ctx, cancle := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控结束")
				return
			default:
				fmt.Println("监控中....")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("1秒后通知监控结束")
	time.Sleep(1 * time.Second)
	cancle()
	time.Sleep(5 * time.Second)
}
