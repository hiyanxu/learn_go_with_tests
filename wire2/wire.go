//+build wireinject

package wire2

import "github.com/google/wire"

func InitA(name string) A {
	wire.Build(NewA)
	return A{}
}
