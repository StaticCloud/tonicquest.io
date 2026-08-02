package orchestrator

import (
	"tonic-quest/globals"
	"tonic-quest/graphics"
	"tonic-quest/sound"
)

type Orchestrator struct {
	Graphics *graphics.Manager
	Sound    *sound.Player
	Globals  *globals.Globals
}

func InitOrchestrator(graphics *graphics.Manager, player *sound.Player, globals *globals.Globals) *Orchestrator {
	return &Orchestrator{
		Graphics: graphics,
		Sound:    player,
		Globals:  globals,
	}
}
