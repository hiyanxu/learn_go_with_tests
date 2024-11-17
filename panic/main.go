package main

import (
	"errors"
	"fmt"

	"github.com/duke-git/lancet/retry"
)

func main() {
	handleError()
}

func handleError() {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case error:
				fmt.Printf("recover err: %+v\n", x)
			default:
				fmt.Printf("recover default: %+v\n", x)
			}
		}
	}()

	_ = retry.Retry(func() (err error) {
		fmt.Println("ddd")
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("recover err: %+v", r)
			}
		}()
		return panicErr()
	})
}

func panicErr() error {
	fmt.Println("ffff")
	panic(errors.New("panic err"))
}
