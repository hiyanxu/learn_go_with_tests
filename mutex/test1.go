package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				// 若没有加锁，则会出现并发读写的问题，导致实际输出的值不是想要的值
				// 主要是由于count++并不是原子操作，需要：先获取值、加1、将结果保存到count变量。
				// 非原子操作可能导致比如10个goroutine同时读取到的值都是100，导致加1操作后都变成了101，本来应该是110.
				//count++

				// 修复1：  采用mutex加互斥锁
				mu.Lock()
				count++
				mu.Unlock()

			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
