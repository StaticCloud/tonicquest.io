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

type Result int

const (
	NO_RESULT Result = iota
	PLAYER_WINS
	ENEMY_WINS
)

type Gameplay struct {
	Orchestrator *orchestrator.Orchestrator
	Player       *entities.Player
	Enemy        *entities.Enemy
	Result       Result
}

func InitGameplayState(orchestrator *orchestrator.Orchestrator) *Gameplay {
	return &Gameplay{
		Orchestrator: orchestrator,
		Player:       entities.InitPlayer(flatKeys, 100),
		Enemy:        entities.InitEnemy([]string{"C", "D", "E"}, 100),
		Result:       NO_RESULT,
	}
}

func (s *Gameplay) Run() (StateName, error) {
	if s.Result == NO_RESULT {
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

		if s.Player.Health <= 0 {
			s.Result = ENEMY_WINS
		}

		if s.Enemy.Health <= 0 {
			s.Result = PLAYER_WINS
		}
	}

	return GAMEPLAY, nil
}

func (s *Gameplay) Draw(screen *ebiten.Image) {
	playerOptions := &ebiten.DrawImageOptions{}

	playerOptions.GeoM.Scale(0.1, 0.1)
	playerOptions.GeoM.Translate(0.0, 70.0)

	screen.DrawImage(
		s.Orchestrator.Graphics.Sprites["carol"],
		playerOptions,
	)

	playerHealthOptions := &ebiten.DrawImageOptions{}

	playerHealthOptions.GeoM.Scale(0.1*float64(s.Player.Health), 0.2)
	playerHealthOptions.GeoM.Translate(0.0, 70.0)

	screen.DrawImage(
		s.Orchestrator.Graphics.Sprites["health"],
		playerHealthOptions,
	)

	enemyOptions := &ebiten.DrawImageOptions{}

	enemyOptions.GeoM.Scale(0.35, 0.35)
	enemyOptions.GeoM.Translate(150.0, 0.0)

	screen.DrawImage(
		s.Orchestrator.Graphics.Sprites["octopussy"],
		enemyOptions,
	)

	enemyHealthOptions := &ebiten.DrawImageOptions{}

	enemyHealthOptions.GeoM.Scale(0.1*float64(s.Enemy.Health), 0.2)
	enemyHealthOptions.GeoM.Translate(150.0, 0.0)

	screen.DrawImage(
		s.Orchestrator.Graphics.Sprites["health"],
		enemyHealthOptions,
	)

	if s.Result == PLAYER_WINS || s.Result == ENEMY_WINS {

	}
	// screen.DrawImage(s.Orchestrator.Graphics.Backgrounds["carol"], &ebiten.DrawImageOptions{})
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
