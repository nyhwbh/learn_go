package app

import (
	"cligame/app/gameinterface"
	"fmt"
)

func GameApp() {
	p := gameinterface.PlayerStatus{}
	m := gameinterface.MonsterStatus{}

	species := 0
	fmt.Scanln(&species)

	p.CreatePlayer()
	p.PlayerSpeciesSetUp(species)
	m.CreateMonster()
}

func UsingSkill(p *gameinterface.PlayerStatus, level int, species int, skill int) {
	switch species {
	case 1:
		switch skill {
		case 1:
			p.Heal()
		case 2:
			p.Steam()
		case 3:
			p.Guard()
		case 4:
			if level >= 99 {
				p.Invincible()
			} else {
				fmt.Println("레벨이 부족하여 사용할 수 없습니다.")
			}
		}
	case 2:
		switch skill {
		case 1:
			p.Heal()
		case 2:
			p.Steam()
		case 3:
			p.Elusion()
		case 4:
			if level >= 99 {
				p.Rapid()
			} else {
				fmt.Println("레벨이 부족하여 사용할 수 없습니다.")
			}
		}
	case 3:
		switch skill {
		case 1:
			p.Heal()
		case 2:
			p.Steam()
		case 3:
			p.Anger()
		case 4:
			if level >= 99 {
				p.Frenzy()
			} else {
				fmt.Println("레벨이 부족하여 사용할 수 없습니다.")
			}
		}
	}
}

func ChangeWeapon(p *gameinterface.PlayerStatus, species int, selection int) {
	switch species {
	case 1:
		switch selection {
		case 1:
			p.Fist()
		case 2:
			p.ShortSword()
		case 3:
			p.LongSword()
		}
	case 2:
		switch selection {
		case 1:
			p.Fist()
		case 2:
			p.ShortBow()
		case 3:
			p.LongBow()
		}
	case 3:
		switch selection {
		case 1:
			p.Fist()
		case 2:
			p.ShortAxe()
		case 3:
			p.IronHammer()
		}
	}
}
