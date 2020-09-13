package main

import (
	"errorGoGenereat/errcode"
	"fmt"
) // 当前的 mod name + package name

func main() {
	var code = errcode.ERR_CODE_INVALID_PARAMS
	fmt.Printf("code: %s\n", code.String())
}