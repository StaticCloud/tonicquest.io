package states

import (
	"tonic-quest/entities/overworld"
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
)

type Overworld struct {
	Orchestrator *orchestrator.Orchestrator
	Player       *overworld.Player
}

func InitOverworldState(orchestrator *orchestrator.Orchestrator) *Overworld {
	return &Overworld{
		Orchestrator: orchestrator,
		Player:       overworld.InitPlayer(),
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
}
