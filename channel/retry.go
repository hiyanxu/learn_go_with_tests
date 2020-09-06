package runner

import (
	"fmt"
	"time"
)

type Consumer struct {
}

var retryChannel chan int

func main() {
	//wg := sync.WaitGroup{}
	retryChannel = make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			retryChannel <- i
		}
	}()

	go Handle()
	time.Sleep(time.Duration(100) * time.Second)
}

func Handle() {
	for i := range retryChannel {
		fmt.Println(i)
	}
}
