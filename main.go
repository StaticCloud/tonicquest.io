package main

import (
	"tonic-quest/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.Init()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tonic Quest")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
