package states

import (
	"tonic-quest/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Gameplay struct {
	Context *utils.Context
}

func InitGameplayState(context *utils.Context) *Gameplay {
	return &Gameplay{
		Context: context,
	}
}

func (s *Gameplay) Run() error {
	return nil
}

func (s *Gameplay) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.Context.Graphics.Backgrounds["carol"], &ebiten.DrawImageOptions{})
}
