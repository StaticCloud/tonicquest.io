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
	Substate         SUBSTATE
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
		Health:     InitHealth(health),
		Coords: &Coords{
			X: 0,
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
		sin := int((e.AnimationManager.Calculate(math.Sin, 120, -0.08, 240, 0)))

		if sin != 120 {
			e.SetX(sin)
		} else {
			e.AnimationManager.Pause = true
		}
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
