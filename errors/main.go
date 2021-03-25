package main

import "fmt"

// 当前的 mod name + package name

//func main() {
//	var code = errcode.ERR_CODE_INVALID_PARAMS
//	fmt.Printf("code: %s\n", code.String())
//}

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func NewError(text string) error {
	// 该处NewError由于没有返回取地址，所以后面比较的时候，仅仅比较值域
	return errorString{s: text}
}

var ErrType = NewError("EOF")

func main() {
	// 由于仅仅比较值域，则相等
	if ErrType == NewError("EOF") {
		fmt.Println("error:", ErrType)
	}
}
