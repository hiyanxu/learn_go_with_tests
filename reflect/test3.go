package reflect

import (
	"fmt"
	"reflect"
)

func IsCanSet() {
	var x float64 = 3.4
	//v := reflect.ValueOf(x)  // 注意 x由于是值，是不能直接调用CanSet的，需要通过取地址取到直接的值，判断CanSet
	v := reflect.ValueOf(&x)
	if v.Elem().CanSet() {
		v.Elem().SetFloat(7.2)
	}

	fmt.Println(x)
}

type SF struct {
	Haha string
	hehe string
}

func GetStructFiled() {
	sf := SF{
		Haha: "hahaha",
		hehe: "hehehe",
	}
	rsfHa := reflect.TypeOf(sf).Field(0)
	fmt.Println("-------hahahah---------")
	fmt.Println(rsfHa.PkgPath)
	fmt.Println("-------hahahah---------")
	fmt.Println("")

	rsfHe := reflect.TypeOf(sf).Field(1)
	fmt.Println("-------hehehehe------")
	fmt.Println(rsfHe.PkgPath)
	fmt.Println("-------hehehehe------")
}
