package reflect

import (
	"reflect"
)

/**
反射：
reflect.ValueOf()：获取值。
reflect.Kind()：获取类型（相关类型有：reflect.Ptr、reflect.Slice、reflect.Struct、reflect.Map、reflect.Array）。
val.NumField()：获取struct中成员变量的数量。
val.Field()：获取所有的字段
val.Len()：获取数组或slice的长度。
val.Index()：获取某个下标的val。
val.MapKeys()：获取map的所有key，返回一个slice
val.MapIndex(key)：获取map的某个key对应的val。
*/

// 反射：运行时检查变量类型，性能较差，难以阅读
func walk(x interface{}, fun func(input string)) {
	//// reflect.ValueOf()返回一个给定变量的value
	//val := reflect.ValueOf(x)
	////field := val.Field(0)
	////fun(field.String())
	//
	//// 若当前类型是指针类型，需要通过Elem()获取指针指向的底层值，相当于(*p)
	//if val.Kind() == reflect.Ptr {
	//	val = val.Elem()
	//}
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value // 将字段的获取过程抽象成一个抽象方法去使用

	switch val.Kind() {
	case reflect.String:
		fun(val.String())
	case reflect.Struct:
		// struct通过NumField获取成员变量数量
		numberOfValues = val.NumField()
		getField = val.Field
		//for i := 0; i < val.NumField(); i++ {
		//	walk(val.Field(i).Interface(), fun)
		//}
	case reflect.Slice:
		// slice通过len获取slice数量
		numberOfValues = val.Len()
		getField = val.Index
	//for j := 0; j < val.Len(); j++ {
	//	walk(val.Index(j).Interface(), fun)
	//}
	case reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fun)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fun)
	}

	// NumField返回值的字段数
	//for i := 0; i < val.NumField(); i++ {
	//	field := val.Field(i)
	//	//fun(field.String())
	//
	//	//// kind()返回值的类型，若是零值，则返回Invalid
	//	//if field.Kind() == reflect.String {
	//	//	fun(field.String())
	//	//}
	//	//
	//	//if field.Kind() == reflect.Struct {
	//	//	// Interface()返回该field的值，并且类型是interface{}
	//	//	walk(field.Interface(), fun)
	//	//}
	//
	//}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
