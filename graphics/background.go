package graphics

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BackgroundManager struct {
	Backgrounds map[string]*ebiten.Image
}

var files map[string]string = map[string]string{
	"splash": "./graphics/sprites/splash.png",
	"carol":  "./graphics/sprites/carol.png",
}

var backgrounds map[string]*ebiten.Image = map[string]*ebiten.Image{}

func InitBackgroundManager() (*BackgroundManager, error) {
	for k, v := range files {
		file, _, err := ebitenutil.NewImageFromFile(v)
		if err != nil {
			return nil, err
		}

		backgrounds[k] = file
	}

	return &BackgroundManager{
		Backgrounds: backgrounds,
	}, nil
}
