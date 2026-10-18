package combat

import (
	"math"
	"tonic-quest/entities/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SubState int

const (
	ENTERING SubState = iota
	IN_COMBAT
	TAKE_DAMAGE
)

type Player struct {
	Moves            []string
	Keys             map[ebiten.Key]string
	AttackList       []string
	Health           *Health
	Coords           *Coords
	Substate         SubState
	AnimationManager *utils.AnimationManager
}

func InitPlayer(keys map[ebiten.Key]string, health int) *Player {
	return &Player{
		Keys:       keys,
		AttackList: []string{},
		Health:     InitHealth(health, 25, 160),
		Coords: &Coords{
			X: 240,
			Y: 25,
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
		p.Entrance()
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

func (p *Player) Entrance() {
	playerX := int((p.AnimationManager.Calculate(math.Sin, 240, -0.05, 240, 1)))
	healthY := int((p.Health.AnimationManager.Calculate(math.Sin, 12, -0.05, 160, 1)))

	if playerX != 0 {
		p.SetX(playerX)
		p.Health.SetY(healthY)
	} else {
		p.Substate = IN_COMBAT
		p.AnimationManager.Pause = true
		p.Health.AnimationManager.Pause = true
	}
}
