package enemy

import (
	"math/rand"
	"time"
	"tonic-quest/keys"
	"tonic-quest/state"
)

type Enemy struct {
	Player     *keys.Player
	Moves      []string
	Manager    *state.Manager
	AttackList []string
}

func InitEnemy(moves []string, player *keys.Player, manager *state.Manager) *Enemy {
	return &Enemy{
		Player:     player,
		Moves:      moves,
		Manager:    manager,
		AttackList: []string{},
	}
}

func (e *Enemy) Attack() {

	for range 3 {
		random := rand.Intn(len(e.Moves))
		time.Sleep(time.Second * 1)
		e.Player.Keys[e.Moves[random]].Rewind()
		e.Player.Keys[e.Moves[random]].Play()
		e.AttackList = append(e.AttackList, e.Moves[random])
	}

	e.Manager.SwitchToPlayerState()
}
