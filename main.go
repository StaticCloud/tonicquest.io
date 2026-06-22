package main

import (
	"tonic-quest/entities"
	"tonic-quest/keys"
	"tonic-quest/state"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game struct {
	GameState *state.Manager
}

var keyBuffer = []ebiten.Key{}

var flatKeys map[ebiten.Key]string = map[ebiten.Key]string{
	ebiten.KeyZ: "C",
	ebiten.KeyX: "D",
	ebiten.KeyC: "E",
	ebiten.KeyV: "F",
	ebiten.KeyB: "G",
	ebiten.KeyN: "A",
	ebiten.KeyM: "B",
}

func (g *Game) Update() error {
	err := g.GameState.Run()

	if err != nil {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	context := audio.NewContext(48000)
	enemy := entities.InitEnemy([]string{"C", "D", "E", "F", "G", "A", "B"}, 100)
	player := entities.InitPlayer(flatKeys, 100)

	soundPlayer, err := keys.InitPlayer(context)
	if err != nil {
		panic(err)
	}

	manager := state.InitStateManager(soundPlayer, player, enemy)

	game := &Game{
		GameState: manager,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tonic Quest")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
