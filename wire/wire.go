//+build wireinject  // wire命令只会处理带有wireinject的文件

package main

import "github.com/google/wire"

func InitMission(name string) Mission {
	wire.Build(NewPlayer, NewMonster, NewMission)
	return Mission{}
}
