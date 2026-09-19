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

func (p *Player) ReadInput(validInputs map[ebiten.Key]bool) {
	keyBuffer := []ebiten.Key{}
	keyBuffer = inpututil.AppendJustPressedKeys(keyBuffer)

	if len(keyBuffer) > 0 {
		key := keyBuffer[0]

		switch key {
		case ebiten.KeyW:
			if validInputs[ebiten.KeyW] == true {
				p.Y -= 20
			}
		case ebiten.KeyS:
			if validInputs[ebiten.KeyS] == true {
				p.Y += 20
			}
		case ebiten.KeyD:
			if validInputs[ebiten.KeyD] == true {
				p.X += 20
			}
		case ebiten.KeyA:
			if validInputs[ebiten.KeyA] == true {
				p.X -= 20
			}
		}
	}
}

func (p *Player) Reset() {
	p.X = 0
	p.Y = 0
}
