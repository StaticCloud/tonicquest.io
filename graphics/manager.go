package graphics

import (
	"bytes"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Manager struct {
	Sprites     map[string]*ebiten.Image
	Backgrounds map[string]*ebiten.Image
	Font        *text.GoTextFace
}

var sourceBackgrounds map[string]string = map[string]string{
	"splash": "./graphics/sprites/splash.png",
	"carol":  "./graphics/sprites/carol.png",
}

var sourceSprites map[string]string = map[string]string{
	"carol":        "./graphics/sprites/carol_sprite.png",
	"health":       "./graphics/sprites/health.png",
	"squelch":      "./graphics/sprites/squelch.png",
	"culo":         "./graphics/sprites/culo.png",
	"twee":         "./graphics/sprites/tree.png",
	"bunkus":       "./graphics/sprites/bunkus.png",
	"michael_cera": "./graphics/sprites/michael_cera.png",
}

var backgrounds map[string]*ebiten.Image = map[string]*ebiten.Image{}
var sprites map[string]*ebiten.Image = map[string]*ebiten.Image{}
var japanese *text.GoTextFaceSource

func InitManager() (*Manager, error) {
	for k, v := range sourceBackgrounds {
		background, _, err := ebitenutil.NewImageFromFile(v)
		if err != nil {
			return nil, err
		}

		backgrounds[k] = background
	}

	for k, v := range sourceSprites {
		sprite, _, err := ebitenutil.NewImageFromFile(v)
		if err != nil {
			return nil, err
		}

		sprites[k] = sprite
	}

	source, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))

	if err != nil {
		log.Fatal(err)
	}

	face := &text.GoTextFace{
		Source: source,
		Size:   11,
	}

	return &Manager{
		Sprites:     sprites,
		Backgrounds: backgrounds,
		Font:        face,
	}, nil
}
