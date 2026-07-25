package overworld

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	X int
	Y int
}

func InitPlayer() *Player {
	return &Player{
		X: 0,
		Y: 0,
	}
}

func (p *Player) ReadInput() {
	keyBuffer := []ebiten.Key{}
	keyBuffer = inpututil.AppendJustPressedKeys(keyBuffer)

	if len(keyBuffer) > 0 {
		key := keyBuffer[0]

		switch key {
		case ebiten.KeyW:
			p.Y -= 20
		case ebiten.KeyS:
			p.Y += 20
		case ebiten.KeyD:
			p.X += 20
		case ebiten.KeyA:
			p.X -= 20
		}
	}
}
