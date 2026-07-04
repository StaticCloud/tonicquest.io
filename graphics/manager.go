package graphics

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Manager struct {
	Backgrounds map[string]*ebiten.Image
}

var files map[string]string = map[string]string{
	"splash": "./graphics/sprites/splash.png",
	"carol":  "./graphics/sprites/carol.png",
}

var backgrounds map[string]*ebiten.Image = map[string]*ebiten.Image{}

func InitManager() (*Manager, error) {
	for k, v := range files {
		file, _, err := ebitenutil.NewImageFromFile(v)
		if err != nil {
			return nil, err
		}

		backgrounds[k] = file
	}

	return &Manager{
		Backgrounds: backgrounds,
	}, nil
}
