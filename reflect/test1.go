package reflect

import (
	"fmt"
	"reflect"
)

type A struct {
	f int64
}

func GetReflectType() {
	a := &A{
		f: 100,
	}
	value := reflect.TypeOf(a)
	fmt.Println(value.Name(), value.String())
}
