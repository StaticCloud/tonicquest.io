package states

import (
	"fmt"
	"math"
	"tonic-quest/entities/combat"
	"tonic-quest/entities/utils"
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
	Player       *combat.Player
	Enemy        *combat.Enemy
	Result       Result
	Timer        *utils.Timer
}

type Load struct {
	Gameplay *Gameplay
}

func InitGameplayState(orchestrator *orchestrator.Orchestrator) *Gameplay {
	return &Gameplay{
		Orchestrator: orchestrator,
		Player:       combat.InitPlayer(flatKeys, 100),
		Result:       NO_RESULT,
		Timer:        utils.InitTimer(),
	}
}

func RefreshCombatState(gameplay *Gameplay) *Load {
	return &Load{
		Gameplay: gameplay,
	}
}

func (l *Load) Run() (StateName, error) {
	l.Gameplay.Enemy = l.Gameplay.Orchestrator.Globals.Enemy
	return COMBAT, nil
}

func (l *Load) Draw(screen *ebiten.Image) {}

func (s *Gameplay) Run() (StateName, error) {
	s.Orchestrator.Graphics.Update()
	s.Timer.Update()
	s.Player.Update()
	s.Enemy.Update()
	s.Player.Health.Entrance(20, 0.3, s.Orchestrator.Globals.Screen.Height, 1)
	s.Enemy.Health.Entrance(20, 0.3, s.Orchestrator.Globals.Screen.Height, 1)

	if s.Result == NO_RESULT {
		key := ""

		if len(s.Enemy.AttackList) == 3 {
			key = s.Player.Attack()

			if len(s.Player.AttackList) > 0 {
				if s.Player.AttackList[len(s.Player.AttackList)-1] != s.Enemy.AttackList[len(s.Player.AttackList)-1] {
					s.Player.TakeDamage(10)
					fmt.Println("Player takes 10 damage!")
					s.Reset()
				}
				if len(s.Player.AttackList) == 3 {
					s.Enemy.TakeDamage(10)
					fmt.Println("Enemy takes 10 damage!")
					s.Reset()
				}

			}
		} else {
			if s.Timer.Pause == false {
				s.Timer.Interval(60)
			} else {
				key = s.Enemy.Attack()
			}
		}

		s.PlayKey(key)

		if s.Player.Health.Value <= 0 {
			s.Player.Health.Reset()

			return OVERWORLD_LOSE, nil
		}

		if s.Enemy.Health.Value <= 0 {
			s.Enemy.Defeated = true
			return OVERWORLD, nil
		}
	}

	return COMBAT, nil
}

func (s *Gameplay) Draw(screen *ebiten.Image) {
	s.DrawHealth(screen, s.Player.Health.Value, (s.Orchestrator.Globals.Screen.Width/2)-110, s.Player.Health.Coords.Y)
	s.DrawHealth(screen, s.Enemy.Health.Value, (s.Orchestrator.Globals.Screen.Width/2)+10, s.Enemy.Health.Coords.Y)

	s.DrawDot(screen)
	s.DrawCarol(screen)
	s.DrawEnemy(screen)
}

func (s *Gameplay) DrawHealth(screen *ebiten.Image, value, x, y int) {
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
	s.Enemy.AttackList = []string{}
	s.Player.AttackList = []string{}
	s.Timer.Reset()
}

func (s *Gameplay) DrawDot(screen *ebiten.Image) {
	for i := range len(s.Orchestrator.Globals.Enemy.AttackList) {
		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(float64((s.Orchestrator.Globals.Screen.Width/2)-10), float64(10+(i*35)))

		screen.DrawImage(
			s.Orchestrator.Graphics.Sprites["note"],
			opts,
		)
	}

	for i := range len(s.Player.AttackList) {
		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(float64((s.Orchestrator.Globals.Screen.Width/2)-10), float64(10+(i*35)))

		screen.DrawImage(
			s.Orchestrator.Graphics.Sprites["note_correct"],
			opts,
		)
	}
}

func (s *Gameplay) CalculateSin(amp, freq float64, offset, limit, ticker int) float64 {
	if math.Sin(float64(ticker)*freq) != float64(limit) {
		return float64(offset) + (amp * math.Sin(float64(ticker)*freq))
	}

	return amp
}

func (s *Gameplay) DrawCarol(screen *ebiten.Image) {
	s.Orchestrator.Graphics.DrawAnimation("carol_idle", s.Player.Coords.X, 10, screen, 0.2)
}

func (s *Gameplay) DrawEnemy(screen *ebiten.Image) {
	s.Orchestrator.Graphics.DrawAnimation(fmt.Sprintf("%s_idle", s.Enemy.Sprite), s.Enemy.Coords.X, 10, screen, 0.2)
}
