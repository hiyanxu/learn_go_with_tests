package main

import (
	"fmt"
	"github.com/dghubble/sling"
)

func main() {
	resp, err := sling.New().Base("https://www.baidu.com").Post("").Request()
	fmt.Println(err, resp)
}
