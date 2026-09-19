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
	Result       Result
}

func InitGameplayState(orchestrator *orchestrator.Orchestrator) *Gameplay {
	return &Gameplay{
		Orchestrator: orchestrator,
		Player:       entities.InitPlayer(flatKeys, 100),
		Result:       NO_RESULT,
	}
}

func (s *Gameplay) Run() (StateName, error) {
	s.Orchestrator.Graphics.Update()
	if s.Result == NO_RESULT {
		key := ""

		if len(s.Orchestrator.Globals.Enemy.AttackList) == 3 {
			key = s.Player.Attack()

			if len(s.Player.AttackList) > 0 {
				if s.Player.AttackList[len(s.Player.AttackList)-1] != s.Orchestrator.Globals.Enemy.AttackList[len(s.Player.AttackList)-1] {
					s.Player.Health -= 10
					fmt.Println("Player takes 10 damage!")
					s.Reset()
				}
				if len(s.Player.AttackList) == 3 {
					s.Orchestrator.Globals.Enemy.Health -= 10
					fmt.Println("Enemy takes 10 damage!")
					s.Reset()
				}
			}
		} else {
			key = s.Orchestrator.Globals.Enemy.Attack()
		}

		s.PlayKey(key)

		if s.Player.Health <= 0 {
			s.Player.Health = 100

			return OVERWORLD_LOSE, nil
		}

		if s.Orchestrator.Globals.Enemy.Health <= 0 {
			s.Orchestrator.Globals.Enemy.Defeated = true
			return OVERWORLD, nil
		}
	}

	return GAMEPLAY, nil
}

func (s *Gameplay) Draw(screen *ebiten.Image) {
	s.drawHealth(screen, s.Player.Health, (s.Orchestrator.Globals.Screen.Width/2)-110, (s.Orchestrator.Globals.Screen.Height)-20)
	s.drawHealth(screen, s.Orchestrator.Globals.Enemy.Health, (s.Orchestrator.Globals.Screen.Width/2)+10, (s.Orchestrator.Globals.Screen.Height)-20)

	s.DrawDot(screen)
}

func (s *Gameplay) drawHealth(screen *ebiten.Image, value, x, y int) {
	borderOptions := &ebiten.DrawImageOptions{}

	borderOptions.GeoM.Translate(float64(x), float64(y))

	screen.DrawImage(
		s.Orchestrator.Graphics.Sprites["health_border"],
		borderOptions,
	)

	healthOptions := &ebiten.DrawImageOptions{}

	healthOptions.GeoM.Scale(float64(value)*0.94, 1.0)
	healthOptions.GeoM.Translate(float64(x+3), float64(y+3))

	screen.DrawImage(
		s.Orchestrator.Graphics.Sprites["health"],
		healthOptions,
	)
}

func (s *Gameplay) PlayKey(key string) {
	if key != "" {
		s.Orchestrator.Sound.Keys[key].Rewind()
		s.Orchestrator.Sound.Keys[key].Play()
	}
}

func (s *Gameplay) Reset() {
	s.Orchestrator.Globals.Enemy.AttackList = []string{}
	s.Player.AttackList = []string{}
}

func (s *Gameplay) DrawDot(screen *ebiten.Image) {
	for i := range len(s.Orchestrator.Globals.Enemy.AttackList) {
		s.Orchestrator.Graphics.DrawAnimation("note", (s.Orchestrator.Globals.Screen.Width / 2), 30+(i*35), screen)
	}
}
