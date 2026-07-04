package orchestrator

import (
	"tonic-quest/graphics"
	"tonic-quest/sound"
)

type Orchestrator struct {
	Graphics *graphics.Manager
	Sound    *sound.Player
}

func InitOrchestrator(graphics *graphics.Manager, player *sound.Player) *Orchestrator {
	return &Orchestrator{
		Graphics: graphics,
		Sound:    player,
	}
}
