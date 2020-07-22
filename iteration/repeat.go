package iteration

import "strings"

func Repeat(character string) string {
	return ""
}

func Repeat1(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		//repeated = repeated + character
		repeated += character
	}

	return repeated
}

func Repeat2(character string, num int) string {
	//var repeated string
	//for i := 0; i < num; i++ {
	//	repeated += character
	//}
	//var repeated []string  // 仅声明未初始化，strings.Join不会生效
	repeated := make([]string, num)
	for i := 0; i < num; i++ {
		repeated = append(repeated, character)
	}

	return strings.Join(repeated, "")
}
