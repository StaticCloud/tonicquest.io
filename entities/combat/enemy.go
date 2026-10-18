package combat

import (
	"math"
	"math/rand"
	"tonic-quest/entities/utils"
)

type Enemy struct {
	ID               int
	Name             string
	Sprite           string
	Moves            []string
	Health           *Health
	Coords           *Coords
	Substate         SubState
	AnimationManager *utils.AnimationManager
	Defeated         bool
	AttackList       []string
	AttackTimer      float64
}

func InitEnemy(ID, health int, name, sprite string, moves []string) *Enemy {
	return &Enemy{
		ID:         ID,
		Name:       name,
		Sprite:     sprite,
		Moves:      moves,
		AttackList: []string{},
		Health:     InitHealth(health, 3, -12),
		Coords: &Coords{
			X: 360,
			Y: 25,
		},
		AnimationManager: utils.InitAnimationManager(),
	}
}

func (e *Enemy) Attack() string {
	e.AttackTimer += 1.0 / 60.0
	attack := ""

	if e.AttackTimer >= 1.0 {
		random := rand.Intn(len(e.Moves))
		attack = e.Moves[random]
		e.AttackList = append(e.AttackList, attack)

		e.AttackTimer = 0
	}

	return attack
}

func (e *Enemy) Update() error {
	if e.AnimationManager.Pause == false {
		e.AnimationManager.Tick += 1
	}
	e.Animate()
	e.Health.Update()
	return nil
}

func (e *Enemy) TakeDamage(amount int) {
	e.Health.TakeDamage(amount)
	e.AnimationManager.Reset()
	e.Substate = TAKE_DAMAGE
}

func (e *Enemy) Reset() {
	e.Health.Reset()
}

func (e *Enemy) SetX(x int) {
	e.Coords.X = x
}

func (e *Enemy) Animate() {
	if e.Substate == ENTERING {
		e.Entrance()
	}

	if e.Substate == TAKE_DAMAGE {
		sin := int((e.AnimationManager.Calculate(math.Sin, 5, 1, 120, 0)))
		if e.AnimationManager.Tick <= 10 {
			e.SetX(sin)
		} else {
			e.AnimationManager.Pause = true
			e.Substate = IN_COMBAT
		}
	}

	if e.Substate == IN_COMBAT {
		e.SetX(120)
	}
}

func (e *Enemy) Entrance() {
	enemyX := int((e.AnimationManager.Calculate(math.Sin, 240, -0.05, 360, 1)))
	healthY := int((e.Health.AnimationManager.Calculate(math.Sin, 24, 0.05, -12, 1)))

	if enemyX != 120 {
		e.SetX(enemyX)
		e.Health.SetY(healthY)
	} else {
		e.Substate = IN_COMBAT
		e.AnimationManager.Pause = true
		e.Health.AnimationManager.Pause = true
	}
}
