package runner

import (
	"fmt"
	"time"
)

// 测试channel初始化和close()，是否会发消息.
func CInit() {
	var closeChennel = make(chan struct{})
	close(closeChennel)

	var c chan struct{}
	//close(c) // c仅声明而未初始化，则close一个nil的channel会panic
	go func() {
		select {
		case _, ok := <-c:
			if !ok {
				fmt.Println("关闭了")
			} else {
				fmt.Println("ddd")
			}
		}
		//for {
		//	select {
		//	case <-c:
		//		fmt.Println("ddd")
		//	}
		//}
	}()

	// 测试init
	if c == nil {
		fmt.Println("nil哦")
		c = closeChennel // 可以直接赋值一个已经close的channel给一个已经声明的channel，这样可以关闭被赋值的channel。例子中的c。
	}
	time.Sleep(time.Second * 1)
}

// 测试close时，是否可以接收到消息.
func CClose() {
	c := make(chan struct{})
	go func() {
		select {
		case _, ok := <-c:
			if !ok {
				fmt.Println("关闭了")
			} else {
				fmt.Println("jjjj")
			}

		}
		//for {
		//	select {
		//	case _, ok := <-c:
		//		if ok {
		//			fmt.Println("ddd")
		//		} else {
		//			fmt.Println("jjjj")
		//		}
		//
		//	}
		//}
	}()

	// 测试init
	//c = make(chan bool)
	close(c)
	time.Sleep(time.Duration(5))
}
