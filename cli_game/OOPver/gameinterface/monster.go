package gameinterface

import (
	"math/rand"
	"time"
)

type Monster struct {
	hp         int
	mp         int
	basicAtk   int
	basicAtks  int
	basicDef   int
	basicAvoid int
}

func (m *Monster) CreateMonset() {
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	m.basicAvoid = 30
	m.hp = random.Intn(50) + 50
	m.mp = random.Intn(50) + 50
	m.basicAtk = random.Intn(10) + 10
	m.basicDef = random.Intn(10) + 5
	m.basicAtks = 800
}
