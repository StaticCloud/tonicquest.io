package states

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type State interface {
	Run() error
	Draw(screen *ebiten.Image)
}
