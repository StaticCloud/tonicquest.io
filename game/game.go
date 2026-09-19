package game

import (
	"tonic-quest/globals"
	"tonic-quest/graphics"
	"tonic-quest/orchestrator"
	"tonic-quest/sound"
	"tonic-quest/states"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	StateManager *states.Manager
	Width        int
	Height       int
}

func Init() *Game {
	graphics, graphicsError := graphics.InitManager()
	player, playerError := sound.InitPlayer()
	globals, globalsError := globals.InitGlobals()

	if graphicsError != nil {
		panic(graphicsError)
	}

	if playerError != nil {
		panic(graphicsError)
	}

	if globalsError != nil {
		panic(globalsError)
	}

	orchestrator := orchestrator.InitOrchestrator(graphics, player, globals)

	return &Game{
		StateManager: states.InitManager(orchestrator),
		Width:        globals.Screen.Width,
		Height:       globals.Screen.Height,
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
	return g.Width, g.Height
}
