package combat

import (
	"math"
	"tonic-quest/entities/utils"
)

type Background struct {
	Coords           *Coords
	AnimationManager *utils.AnimationManager
	Substate         SubState
}

func InitBackground() *Background {
	return &Background{
		Coords: &Coords{
			X: 240,
			Y: 20,
		},
		Substate:         ENTERING,
		AnimationManager: utils.InitAnimationManager(),
	}
}

func (b *Background) Update() error {
	if b.AnimationManager.Pause == false {
		b.AnimationManager.Tick += 1
	}
	b.Animate()
	return nil
}

func (b *Background) SetX(x int) {
	b.Coords.X = x
}

func (b *Background) Animate() {
	if b.Substate == ENTERING {
		sin := int(b.AnimationManager.Calculate(math.Sin, 300, -0.08, 240, 0))
		if sin >= -55 {
			b.SetX(sin)
		} else {
			b.AnimationManager.Pause = true
			b.Substate = IN_COMBAT
		}
	}

	if b.Substate == IN_COMBAT {
		b.SetX(-60)
	}
}
