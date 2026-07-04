package game

import (
	"tonic-quest/graphics"
	"tonic-quest/orchestrator"
	"tonic-quest/sound"
	"tonic-quest/states"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	StateManager *states.Manager
}

func Init() *Game {
	graphics, graphicsError := graphics.InitManager()
	player, playerError := sound.InitPlayer()

	if graphicsError != nil {
		panic(graphicsError)
	}

	if playerError != nil {
		panic(graphicsError)
	}

	orchestrator := orchestrator.InitOrchestrator(graphics, player)

	return &Game{
		StateManager: states.InitManager(orchestrator),
	}
}

func (g *Game) Update() error {
	if err := g.StateManager.Run(); err != nil {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.StateManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 256, 224
}
