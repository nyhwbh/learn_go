package player

import (
	"math/rand"
	"time"
)

type PlayerStatus struct {
	level       int
	hp          int
	maxhp       int
	mp          int
	maxmp       int
	basicAtk    int
	battleAtk   int
	basicAtks   int
	battleAtks  int
	basicDef    int
	battleDef   int
	basicAvoid  int
	battleAvoid int
	species     int
}

// 캐릭터 생성
func (p *PlayerStatus) CreatePlayer() {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	p.level = 1
	p.basicAvoid = 30
	p.battleAvoid = p.basicAvoid
	p.maxhp = random.Intn(50) + 50
	p.hp = p.maxhp
	p.maxmp = random.Intn(50) + 50
	p.mp = p.maxmp
	p.basicAtk = random.Intn(10) + 10
	p.battleAtk = p.basicAtk
	p.basicDef = random.Intn(10) + 5
	p.battleDef = p.basicDef
	p.basicAtks = 1000
	p.battleAtks = p.basicAtks
}

// 캐릭터 종족 선택
func (p *PlayerStatus) PlayerSpeciesSetUp(sp int) {
	p.species = sp
}

// 캐릭터 레벨업
func (p *PlayerStatus) PlayerLevelUp(sp int) {
	p.level += 1

	// 최대체력 증가
	if p.maxhp == 150 {
	} else {
		if p.maxhp+3 > 150 {
			p.maxhp = 150
		} else {
			p.maxhp += 3
		}
	}

	// 최대 마력 증가
	if p.maxmp == 150 {
	} else {
		if p.maxmp+3 > 150 {
			p.maxmp = 150
		} else {
			p.maxmp += 3
		}
	}
	// 종족별 능력치 상승
	if p.level%5 == 0 {
		switch sp {
		case 1: // Human
			p.StatusLevelUpHuman()
		case 2: // Elf
			p.StatusLevelUpElf()
		case 3: // Oak
			p.StatusLevelUpOak()
		}
	}
}

// 종족별 레벨업 능력치
// human
func (p *PlayerStatus) StatusLevelUpHuman() {
	p.basicAtk += 1
	p.basicDef += 1
}

// elf
func (p *PlayerStatus) StatusLevelUpElf() {
	p.basicAtk += 1
	p.basicAtks += 100
}

// oak
func (p *PlayerStatus) StatusLevelUpOak() {
	p.basicAtk += 2
}

// 캐릭터 휴식 -> 체력 및 마나 회복
func (p *PlayerStatus) PlayerRecovery() {
	p.hp = p.maxhp
	p.mp = p.maxmp
}
