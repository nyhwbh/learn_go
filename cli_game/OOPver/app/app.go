package app

import (
	"cligame/app/gameinterface"
	"fmt"
)

func app() {
	p := gameinterface.PlayerStatus{}
	m := gameinterface.MonsterStatus{}

	species := 0
	fmt.Scanln(&species)

	p.CreatePlayer()
	p.PlayerSpeciesSetUp(species)
	m.CreateMonster()
}
