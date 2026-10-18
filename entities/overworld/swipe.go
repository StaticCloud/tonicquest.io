package overworld

import (
	"math"
	"tonic-quest/entities/utils"
)

type Swipe struct {
	X                int
	AnimationManager *utils.AnimationManager
}

func InitSwipe() *Swipe {
	return &Swipe{
		X:                240,
		AnimationManager: utils.InitAnimationManager(),
	}
}

func (s *Swipe) Update() error {
	if s.AnimationManager.Pause == false {
		s.AnimationManager.Tick += 1
	}
	return nil
}

func (s *Swipe) Start() {
	s.AnimationManager.Pause = false
	s.TransitionOut()
}

func (s *Swipe) SetX(x int) {
	s.X = x
}

func (s *Swipe) TransitionOut() {
	sin := int((s.AnimationManager.Calculate(math.Sin, 300, -0.08, 240, 1)))
	if sin != 0 {
		s.SetX(sin)
	}
}
