package main

import (
	"fmt"
)

type Monster struct {
	Name string
}

func NewMonster(name string) Monster {
	return Monster{Name: name}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(player Player, monster Monster) Mission {
	return Mission{
		Player:  player,
		Monster: monster,
	}
}

func (m *Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

func main() {
	mission := InitMission("dj")
	mission.Start()
}
