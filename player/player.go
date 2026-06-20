package player

import (
	"tonic-quest/keys"
	"tonic-quest/state"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Player     *keys.Player
	Moves      []string
	Keys       map[ebiten.Key]string
	Manager    *state.Manager
	AttackList []string
}

func InitPlayer(player *keys.Player, keys map[ebiten.Key]string, manager *state.Manager) *Player {
	return &Player{
		Player:     player,
		Keys:       keys,
		Manager:    manager,
		AttackList: []string{},
	}
}

func (p *Player) GetMove() string {
	keyBuffer := []ebiten.Key{}
	keyBuffer = inpututil.AppendJustPressedKeys(keyBuffer)

	if len(keyBuffer) > 0 {
		input := keyBuffer[0]
		key := p.Keys[input]

		if key != "" {
			p.Player.Keys[key].Rewind()
			p.Player.Keys[key].Play()
		}

		return key
	}

	return ""
}
