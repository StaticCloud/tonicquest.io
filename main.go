package main

import (
	"tonic-quest/entities"
	"tonic-quest/game"
	"tonic-quest/graphics"
	"tonic-quest/keys"
	"tonic-quest/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game struct {
	GameManager *game.Manager
	Graphics    *graphics.BackgroundManager
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
	err := g.GameManager.Run()

	if err != nil {
		return ebiten.Termination
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.GameManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 256, 224
}

func main() {
	context := audio.NewContext(48000)
	enemy := entities.InitEnemy([]string{"C", "D", "E", "F", "G", "A", "B"}, 100)
	player := entities.InitPlayer(flatKeys, 100)

	background, backgroundErr := graphics.InitBackgroundManager()
	if backgroundErr != nil {
		panic(backgroundErr)
	}

	soundPlayer, err := keys.InitPlayer(context)
	if err != nil {
		panic(err)
	}

	gameContext := &utils.Context{
		Player:   player,
		Enemy:    enemy,
		Sound:    soundPlayer,
		Graphics: background,
	}
	manager := game.InitGameManager(gameContext)

	game := &Game{
		GameManager: manager,
		Graphics:    background,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tonic Quest")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
