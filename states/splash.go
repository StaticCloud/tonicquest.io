package states

import (
	"time"
	"tonic-quest/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Splash struct {
	Context *utils.Context
}

func InitSplashState(context *utils.Context) *Splash {
	return &Splash{
		Context: context,
	}
}

func (s *Splash) Run() error {
	time.Sleep(time.Second * 3)

	return nil
}

func (s *Splash) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.Context.Graphics.Backgrounds["splash"], &ebiten.DrawImageOptions{})
}
