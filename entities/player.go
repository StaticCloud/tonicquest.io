package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Moves      []string
	Keys       map[ebiten.Key]string
	AttackList []string
	Health     int
}

func InitPlayer(keys map[ebiten.Key]string, health int) *Player {
	return &Player{
		Keys:       keys,
		AttackList: []string{},
		Health:     health,
	}
}

func (p *Player) Attack() string {
	keyBuffer := []ebiten.Key{}
	keyBuffer = inpututil.AppendJustPressedKeys(keyBuffer)

	if len(keyBuffer) > 0 {
		return p.Keys[keyBuffer[0]]
	}
	return ""
}
