package maps

import "fmt"

type A struct {
	m1 map[string]interface{}
}

func HandleReadWriteMap() {
	a := &A{
		m1: map[string]interface{}{
			"f1": 1,
			"f2": "2",
		},
	}

	a1 := *a
	a1.m1["f1"] = 11

	fmt.Println(a)
	fmt.Println(a1)

	fmt.Println("test git delete")

	a1.m1 = make(map[string]interface{}) // 重新赋一个新的 map，不改原来的.
	fmt.Println(a)
	fmt.Println(a1)

	fmt.Println("Hello, 世界")
}
