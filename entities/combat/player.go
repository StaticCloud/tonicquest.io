package combat

import (
	"math"
	"tonic-quest/entities/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SUBSTATE int

const (
	ENTERING = iota
	IN_COMBAT
	TAKE_DAMAGE
)

type Player struct {
	Moves            []string
	Keys             map[ebiten.Key]string
	AttackList       []string
	Health           *Health
	Coords           *Coords
	Substate         SUBSTATE
	AnimationManager *utils.AnimationManager
}

func InitPlayer(keys map[ebiten.Key]string, health int) *Player {
	return &Player{
		Keys:       keys,
		AttackList: []string{},
		Health:     InitHealth(health),
		Coords: &Coords{
			X: -120,
		},
		Substate:         ENTERING,
		AnimationManager: utils.InitAnimationManager(),
	}
}

func (p *Player) Attack() string {
	keyBuffer := []ebiten.Key{}
	keyBuffer = inpututil.AppendJustPressedKeys(keyBuffer)

	if len(keyBuffer) > 0 {
		key := p.Keys[keyBuffer[0]]
		p.AttackList = append(p.AttackList, key)
		return key
	}
	return ""
}

func (p *Player) Update() error {
	if p.AnimationManager.Pause == false {
		p.AnimationManager.Tick += 1
	}
	p.Animate()
	p.Health.Update()
	return nil
}

func (p *Player) TakeDamage(amount int) {
	p.Health.TakeDamage(amount)
	p.AnimationManager.Reset()
	p.Substate = TAKE_DAMAGE
}

func (p *Player) Reset() {
	p.Health.Reset()
}

func (p *Player) SetX(x int) {
	p.Coords.X = x
}

func (p *Player) Animate() {
	if p.Substate == ENTERING {
		sin := int((p.AnimationManager.Calculate(math.Sin, 120, 0.08, -120, 1)))
		if sin != 0 {
			p.SetX(sin)
		} else {
			p.Substate = IN_COMBAT
			p.AnimationManager.Pause = true
		}
	}

	if p.Substate == TAKE_DAMAGE {
		sin := int((p.AnimationManager.Calculate(math.Sin, 5, 1, 0, 0)))
		if p.AnimationManager.Tick <= 10 {
			p.SetX(sin)
		} else {
			p.AnimationManager.Pause = true
			p.Substate = IN_COMBAT
		}
	}

	if p.Substate == IN_COMBAT {
		p.SetX(0)
	}
}
