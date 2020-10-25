package reflect

import (
	"reflect"
)

type MyInt int

type S struct {
	Str string
}

func getKind() []reflect.Kind {
	var i1 MyInt = 100
	iKind := reflect.ValueOf(i1).Type().Kind()
	s1 := &S{Str: "str1"}
	sKind := reflect.ValueOf(s1).Type().Kind()
	s2 := S{Str: "str2"}
	s2Kind := reflect.ValueOf(s2).Type().Kind()
	str1 := "ddd"
	strKind := reflect.ValueOf(str1).Type().Kind()
	//fmt.Println(iKind)
	//fmt.Println(sKind)
	//fmt.Println(strKind)
	return []reflect.Kind{iKind, sKind, s2Kind, strKind}
}
