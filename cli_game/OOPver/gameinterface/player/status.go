package player

import (
	"math/rand"
	"time"
)

type PlayerStatus struct {
	level   int
	hp      int
	maxhp   int
	mp      int
	maxmp   int
	atk     int
	atks    int
	def     int
	avoid   int
	species int
}

func (p *PlayerStatus) BasicSetup() {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	p.level = 1
	p.avoid = 30
	p.maxhp = random.Intn(50) + 50
	p.hp = p.maxhp
	p.maxmp = random.Intn(50) + 50
	p.mp = p.maxmp
	p.atk = random.Intn(10) + 20
	p.def = random.Intn(10) + 10
	p.atks = 1000
}

func (p PlayerStatus) PlayerSpeciesSetUp(sp int) {
	p.species = sp
}

func (p *PlayerStatus) PlayerLevelUp() {
	p.level += 1
}
