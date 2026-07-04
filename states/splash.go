package states

import (
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
)

type Splash struct {
	Orchestrator *orchestrator.Orchestrator
	timer        float64
}

func InitSplashState(orchestrator *orchestrator.Orchestrator) *Splash {
	return &Splash{
		Orchestrator: orchestrator,
	}
}

func (s *Splash) Run() (StateName, error) {
	s.timer += 1.0 / 60.0 // increment by one tick
	if s.timer >= 3.0 {
		return GAMEPLAY, nil
	}
	return SPLASH, nil
}

func (s *Splash) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.Orchestrator.Graphics.Backgrounds["splash"], &ebiten.DrawImageOptions{})
}
