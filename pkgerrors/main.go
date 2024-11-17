package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := func3()
	fmt.Printf("%+v", err)
}

func func1() error {
	return fmt.Errorf("func1Err")
}

func func2() error {
	err := func1()
	return errors.WithStack(err)
}

func func3() error {
	err := func2()
	/**
	当func2已经用withStack包装过一次之后，func3再用withStack包装，则会将error重复2遍
	输出如下：
	func1Err
	main.func2
		/Users/bjhl/Documents/baijiahulian/learn_go_with_tests/pkgerrors/example.go:19
	main.func3
		/Users/bjhl/Documents/baijiahulian/learn_go_with_tests/pkgerrors/example.go:23
	main.main
		/Users/bjhl/Documents/baijiahulian/learn_go_with_tests/pkgerrors/example.go:9
	runtime.main
		/usr/local/go/src/runtime/proc.go:203
	runtime.goexit
		/usr/local/go/src/runtime/asm_amd64.s:1357
	main.func3
		/Users/bjhl/Documents/baijiahulian/learn_go_with_tests/pkgerrors/example.go:24
	main.main
		/Users/bjhl/Documents/baijiahulian/learn_go_with_tests/pkgerrors/example.go:9
	runtime.main
		/usr/local/go/src/runtime/proc.go:203
	runtime.goexit
		/usr/local/go/src/runtime/asm_amd64.s:1357%

	所以，建议不要重复包装
	使用withStack的场景：
	1、调用github等第三方包。
	2、调用底层包。
	3、调用第三方接口。

	错误堆栈重复的原因：
	func WithStack(err error) error {
		if err == nil {
			return nil
		}
		return &withStack{
			err,
			callers(),
		}
	}
	如上所示，返回withStack时，会调用callers()，该函数会重复获取调用堆栈。
	*/
	return errors.WithStack(err)
}
