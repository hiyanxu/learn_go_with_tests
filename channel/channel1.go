package runner

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

type a struct {
	ready chan bool
}

func (a *a) work() error {
	time.Sleep(3 * time.Second)
	//return errors.New("err")
	return nil
}

func Start() error {
	a1 := &a{ready: make(chan bool)}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("开始执行defer...")

		for {
			if err := a1.work(); err != nil {
				fmt.Printf("出现error: %s\n", err)
			}

			a1.ready <- true

		}
	}()

	<-a1.ready
	fmt.Printf("start and running...")

	return nil
}
