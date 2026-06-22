package entities

import (
	"math/rand"
	"time"
)

type Enemy struct {
	Moves      []string
	Health     int
	AttackList []string
}

func InitEnemy(moves []string, health int) *Enemy {
	return &Enemy{
		Moves:      moves,
		Health:     health,
		AttackList: []string{},
	}
}

func (e *Enemy) Attack() string {
	random := rand.Intn(len(e.Moves))
	time.Sleep(time.Second * 1)
	attack := e.Moves[random]
	e.AttackList = append(e.AttackList, attack)

	return attack
}

func (e *Enemy) TakeDamage(amount int) {
	e.Health = e.Health - amount
}
