package struct_interface

import "math"

func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func Area(width, height float64) float64 {
	return width * height
}

/**
方法：
方法：调用方法时，数据的引用是通过receiverName获取，其它语言中通过隐式的this变量.
惯例：将类型的第一个字母作为接收者变量时Go语言的惯例。
*/
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func PerimeterForStruct(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

func AreaForStruct(r Rectangle) float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
