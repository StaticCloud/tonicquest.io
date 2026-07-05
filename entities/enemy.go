package entities

import (
	"math/rand"
)

type Enemy struct {
	Moves      []string
	Health     int
	AttackList []string
	timer      float64
}

func InitEnemy(moves []string, health int) *Enemy {
	return &Enemy{
		Moves:      moves,
		Health:     health,
		AttackList: []string{},
	}
}

func (e *Enemy) Attack() string {
	e.timer += 1.0 / 60.0
	attack := ""

	if e.timer >= 1.0 {
		random := rand.Intn(len(e.Moves))
		attack = e.Moves[random]
		e.AttackList = append(e.AttackList, attack)

		e.timer = 0
	}

	return attack
}

func (e *Enemy) TakeDamage(amount int) {
	e.Health = e.Health - amount
}
