package main

import (
	"log"
	"tonic-quest/keys"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Player *keys.Player
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
	keyBuffer = inpututil.AppendJustPressedKeys(keyBuffer)
	if len(keyBuffer) > 0 {
		input := keyBuffer[0]
		key := flatKeys[input]

		if key != "" {
			g.Player.Keys[key].Rewind()
			g.Player.Keys[key].Play()
		}
	}
	keyBuffer = []ebiten.Key{}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	context := audio.NewContext(48000)

	player, err := keys.InitPlayer(context)
	if err != nil {
		panic(err)
	}

	game := &Game{
		Player: player,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tonic Quest")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
