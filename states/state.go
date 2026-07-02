package states

import (
	"tonic-quest/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type State interface {
	Run() error
	Draw(screen *ebiten.Image)
	SetEmitter(emitter utils.Emitter)
}
