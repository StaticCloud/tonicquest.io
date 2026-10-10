package combat

import (
	"math"
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

func InitHealth(health int) *Health {
	return &Health{
		Initial:          health,
		Value:            health,
		Coords:           &Coords{},
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

func (h *Health) Entrance(amp, speed float64, offset, limit int) {
	sin := int((h.AnimationManager.Calculate(math.Sin, amp, speed, offset, limit)))
	if sin != 140 {
		h.SetY(sin)
	} else {
		h.AnimationManager.Pause = true
	}
}
