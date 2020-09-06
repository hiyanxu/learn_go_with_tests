package runner

import (
	"fmt"
	"testing"
	"time"
)

const timeout = 3 * time.Second

func TestNewRunner(t *testing.T) {
	fmt.Println("start.........")

	r := NewRunner(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOunt:
			fmt.Printf("err timout \n")
		case ErrInterrupt:
			fmt.Printf("err interrupt\n")
		}
	}

	fmt.Printf("complete\n")
}

func createTask() func(int) {
	return func(i int) {
		time.Sleep(2 * time.Second)
		fmt.Printf("id: %d\n", i)
	}
}
