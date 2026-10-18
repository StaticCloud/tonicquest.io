package combat

import (
	"tonic-quest/entities/utils"
)

type Coords struct {
	X int
	Y int
}

type Health struct {
	Initial          int
	Value            int
	Coords           *Coords
	AnimationManager *utils.AnimationManager
}

func InitHealth(health, x, y int) *Health {
	return &Health{
		Initial: health,
		Value:   health,
		Coords: &Coords{
			X: x,
			Y: y,
		},
		AnimationManager: utils.InitAnimationManager(),
	}
}

func (h *Health) TakeDamage(amount int) {
	h.Value -= amount
}

func (h *Health) Reset() {
	h.Value = h.Initial
}

func (h *Health) Update() error {
	if h.AnimationManager.Pause == false {
		h.AnimationManager.Tick += 1
	}
	return nil
}

func (h *Health) SetY(y int) {
	h.Coords.Y = y
}

func (h *Health) SetX(x int) {
	h.Coords.X = x
}
