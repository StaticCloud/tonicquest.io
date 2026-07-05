package states

import (
	"fmt"
	"tonic-quest/entities"
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
)

var flatKeys map[ebiten.Key]string = map[ebiten.Key]string{
	ebiten.KeyZ: "C",
	ebiten.KeyX: "D",
	ebiten.KeyC: "E",
	ebiten.KeyV: "F",
	ebiten.KeyB: "G",
	ebiten.KeyN: "A",
	ebiten.KeyM: "B",
}

type Gameplay struct {
	Orchestrator *orchestrator.Orchestrator
	Player       *entities.Player
	Enemy        *entities.Enemy
}

func InitGameplayState(orchestrator *orchestrator.Orchestrator) *Gameplay {
	return &Gameplay{
		Orchestrator: orchestrator,
		Player:       entities.InitPlayer(flatKeys, 100),
		Enemy:        entities.InitEnemy([]string{"C", "D", "E"}, 100),
	}
}

func (s *Gameplay) Run() (StateName, error) {
	key := ""

	if len(s.Enemy.AttackList) == 3 {
		key = s.Player.Attack()

		if len(s.Player.AttackList) > 0 {
			if s.Player.AttackList[len(s.Player.AttackList)-1] != s.Enemy.AttackList[len(s.Player.AttackList)-1] {
				s.Player.Health -= 10
				fmt.Println("Player takes 10 damage!")
				s.Reset()
			}
			if len(s.Player.AttackList) == 3 {
				s.Enemy.Health -= 10
				fmt.Println("Enemy takes 10 damage!")
				s.Reset()
			}
		}
	} else {
		key = s.Enemy.Attack()
	}

	s.PlayKey(key)

	return GAMEPLAY, nil
}

func (s *Gameplay) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.Orchestrator.Graphics.Backgrounds["carol"], &ebiten.DrawImageOptions{})
}

func (s *Gameplay) PlayKey(key string) {
	if key != "" {
		s.Orchestrator.Sound.Keys[key].Rewind()
		s.Orchestrator.Sound.Keys[key].Play()
	}
}

func (s *Gameplay) Reset() {
	s.Enemy.AttackList = []string{}
	s.Player.AttackList = []string{}
}
