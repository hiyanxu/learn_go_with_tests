package _select

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) string {
	//startA := time.Now()
	//http.Get(a)
	//// 返回从某个时间到当前时间的间隔
	//aDuration := time.Since(startA)
	//
	//startB := time.Now()
	//http.Get(b)
	//bDuration := time.Since(startB)
	aDuration := measureResponse(a)
	bDuration := measureResponse(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponse(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer2(a, b string) (winner string) {
	// select 允许同时在 多个channel 等待，哪个channel获取值，则哪个case执行
	// range的区别：range会一直阻塞，直到channel有数据时，执行 例如：for c := range cha1 { fmt.Println(c) }
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}

const TimedOut = 10

// 正常测试时，使用Racer3
func Racer3(a, b string) (string, error) {
	return ConfigurableRacer(a, b, TimedOut)
}

// 超时情况的测试，可以使用该函数，自定义超时时间
func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout * time.Second):
		return "", fmt.Errorf("timed out waiting for %s, and %s", a, b)
	}
}
