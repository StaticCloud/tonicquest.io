package utils

type AnimationManager struct {
	Tick  int
	Pause bool
}

func InitAnimationManager() *AnimationManager {
	return &AnimationManager{
		Tick:  0,
		Pause: false,
	}
}

func (a *AnimationManager) Reset() {
	a.Tick = 0
	a.Pause = false
}

func (a *AnimationManager) Calculate(f func(float64) float64, amp, freq float64, offset, limit int) float64 {
	return float64(offset) + (amp * f(float64(a.Tick)*freq))
}
