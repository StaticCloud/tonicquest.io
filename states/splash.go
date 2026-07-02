package states

import (
	"tonic-quest/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Splash struct {
	Context *utils.Context
	Emitter utils.Emitter
	timer   float64
}

func InitSplashState(context *utils.Context) *Splash {
	return &Splash{
		Context: context,
	}
}

func (s *Splash) Run() error {
	s.timer += 1.0 / 60.0 // increment by one tick
	if s.timer >= 3.0 {
		s.Emitter.Emit(utils.SWITCH_TO_GAMEPLAY)
	}
	return nil
}

func (s *Splash) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.Context.Graphics.Backgrounds["splash"], &ebiten.DrawImageOptions{})
}

func (s *Splash) SetEmitter(emitter utils.Emitter) {
	s.Emitter = emitter
}
