package main

import (
	"fmt"
)

/**
重构过程：
1、提取常量。
2、if分支增多时，使用switch代替if。
3、switch可以单独提取方法。
 */

const englishHelloPrefix  = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "

func Hello() string  {
	return "hello, world"
}

func Hello2(name string) string {
	return "Hello, " + name
}

func Hello3(name string) string {
	return englishHelloPrefix + name
}

//func Hello4(name string, language string) string {
//	if name == "" {
//		name = "World"
//	}
//	if language == spanish {
//		return spanishHelloPrefix + name
//	}
//
//	return englishHelloPrefix + name
//}

//func Hello4(name string, language string) string  {
//	if name == "" {
//		name = "World"
//	}
//	switch language {
//	case french:
//		return frenchHelloPrefix + name
//	case spanish:
//		return spanishHelloPrefix + name
//	default:
//		return englishHelloPrefix + name
//	}
//}

func Hello4(name string, language string) string  {
	if name == "" {
		name = "World"
	}

	prefix := "Hello, "
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func Hello5(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main()  {
	fmt.Println(Hello2("world"))
}
