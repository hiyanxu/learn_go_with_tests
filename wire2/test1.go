package wire2

import (
	"github.com/dghubble/sling"
)

type A struct {
	name string
	s    *sling.Sling
}

func NewA(name string) A {
	return A{
		name: name,
		s:    sling.New(),
	}
}
