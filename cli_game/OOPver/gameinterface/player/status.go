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

func (c *PlayerStatus) BasicSetup() {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	c.level = 1
	c.avoid = 30
	c.maxhp = random.Intn(50) + 50
	c.hp = c.maxhp
	c.maxmp = random.Intn(50) + 50
	c.mp = c.maxmp
	c.atk = random.Intn(10) + 20
	c.def = random.Intn(10) + 10
	c.atks = 1000
}
