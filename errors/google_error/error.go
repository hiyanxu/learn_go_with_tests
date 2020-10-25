package google_error

import (
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type retryInfo = errdetails.RetryInfo

/**
别名类型：两者除了名字外完全相同，可以比较，可以相互赋值
类型再定义：两个不同的类型，两个类型的变量不可以相互赋值，也不可以判断或比较。类型再定义的两个类型却可以进行【类型转换】

另一个重要区别：
类型再定义的类型的方法集合原始类型没有任何关系。
别名类型和原始类型的方法集是相同的。
*/

// 别名类型：两者除了名字外完全相同，可以比较，可以相互赋值
type myString = string

func haha() {
	var ms1 myString = "ms1"
	var s1 string = "ms2"
	ms1 = s1

	fmt.Println(ms1)
}

// 类型再定义：两个不同的类型，两个类型的变量不可以相互赋值，也不可以判断或比较
// 类型再定义的两个类型却可以进行【类型转换】
type myString2 string

func hehe() {
	var ms1 myString2 = "ms1"
	var s1 string = "ms2"
	ms1 = s1
}
