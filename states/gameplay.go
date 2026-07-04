package states

import (
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
}

func InitGameplayState(orchestrator *orchestrator.Orchestrator) *Gameplay {
	return &Gameplay{
		Orchestrator: orchestrator,
		Player:       entities.InitPlayer(flatKeys, 100),
	}
}

func (s *Gameplay) Run() (StateName, error) {

	if key := s.Player.Attack(); key != "" {
		s.Orchestrator.Sound.Keys[key].Rewind()
		s.Orchestrator.Sound.Keys[key].Play()
	}

	return GAMEPLAY, nil
}

func (s *Gameplay) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.Orchestrator.Graphics.Backgrounds["carol"], &ebiten.DrawImageOptions{})
}
