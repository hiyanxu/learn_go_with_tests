package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	Count int
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.Count++
	c.mu.Unlock()
}

func (c *Counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Count
}

func Test() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.GetCount())
}
