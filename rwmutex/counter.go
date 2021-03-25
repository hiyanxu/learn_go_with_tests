package rwmutex

import "sync"

// Counter 计数器.
type Counter struct {
	Count   int
	RWMutex sync.RWMutex
}

func (c *Counter) Incr() {
	// 写操作使用Lock和Unlock
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	c.Count++
}

func (c *Counter) Get() int {
	// 读操作使用RLock和RUnlock，提高并发读写性能
	c.RWMutex.RLock()
	c.RWMutex.RUnlock()
	return c.Count
}
