package states

import (
	"tonic-quest/entities/overworld"
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
)

type Overworld struct {
	Orchestrator *orchestrator.Orchestrator
	Player       *overworld.Player
	Enemies      []*overworld.Enemy
}

func InitOverworldState(orchestrator *orchestrator.Orchestrator) *Overworld {
	return &Overworld{
		Orchestrator: orchestrator,
		Player:       overworld.InitPlayer(),
		Enemies: []*overworld.Enemy{
			overworld.InitEnemy(200, 200),
		},
	}
}

func (s *Overworld) Run() (StateName, error) {
	s.Player.ReadInput()
	return OVERWORLD, nil
}

func (s *Overworld) Draw(screen *ebiten.Image) {
	playerOptions := &ebiten.DrawImageOptions{}

	playerOptions.GeoM.Scale(0.1, 0.1)
	playerOptions.GeoM.Translate(float64(s.Player.X), float64(s.Player.Y))

	screen.DrawImage(
		s.Orchestrator.Graphics.Sprites["carol"],
		playerOptions,
	)

	for _, enemy := range s.Enemies {
		enemyOptions := &ebiten.DrawImageOptions{}

		enemyOptions.GeoM.Scale(0.1, 0.1)
		enemyOptions.GeoM.Translate(float64(enemy.X), float64(enemy.Y))

		screen.DrawImage(
			s.Orchestrator.Graphics.Sprites["octopussy"],
			enemyOptions,
		)
	}
}
