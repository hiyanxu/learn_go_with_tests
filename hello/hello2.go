// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type AI interface {
	Haha(string)
}

type A struct{}

func (a *A) Haha(name string) {
	fmt.Println(name)
}

func (a A) Hehe() {
	fmt.Println("hehe")
}

func main() {
	var ai AI

	a := &A{}
	a.Haha("h1")
	a.Hehe()
	ai = a
	ai.Haha("hhhh")

	b := A{}
	ai = b
	b.Hehe()
	b.Haha("h2")

	fmt.Println("Hello, 世界")
}
