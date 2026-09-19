package graphics

import (
	"bytes"
	"image"
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
	Timer       int64
	Ticker      int
}

type Animation struct {
	SpriteSheet *ebiten.Image
	Frame0X     int
	Frame0Y     int
	FrameWidth  int
	FrameHeight int
	FrameCount  int
}

var sourceBackgrounds map[string]string = map[string]string{
	"splash": "./graphics/sprites/splash.png",
	"carol":  "./graphics/sprites/carol.png",
}

var sourceSprites map[string]string = map[string]string{
	"carol":         "./graphics/sprites/carol_sprite.png",
	"health":        "./graphics/sprites/combat/health.png",
	"health_border": "./graphics/sprites/combat/health-border.png",
	"squelch":       "./graphics/sprites/squelch.png",
	"culo":          "./graphics/sprites/culo.png",
	"twee":          "./graphics/sprites/tree.png",
	"bunkus":        "./graphics/sprites/bunkus.png",
	"michael_cera":  "./graphics/sprites/michael_cera.png",
	"typewriter":    "./graphics/sprites/typewriter.png",
}

var animatedSprites map[string]*Animation = map[string]*Animation{
	"note": InitAnimation("./graphics/sprites/note_indicator.png", 20, 20, 5),
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

func InitAnimation(source string, width, height, count int) *Animation {
	sprite, _, _ := ebitenutil.NewImageFromFile(source)

	return &Animation{
		SpriteSheet: sprite,
		Frame0X:     0,
		Frame0Y:     0,
		FrameWidth:  width,
		FrameHeight: height,
		FrameCount:  count,
	}
}

func (m *Manager) Update() error {
	m.Ticker++
	return nil
}

func (m *Manager) DrawAnimation(sprite string, x, y int, screen *ebiten.Image) {
	animation := animatedSprites[sprite]
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(animation.FrameWidth)/2, -float64(animation.FrameHeight)/2)
	op.GeoM.Translate(float64(x), float64(y))
	i := (m.Ticker / 5) % animation.FrameCount
	sx, sy := animation.Frame0X+i*animation.FrameWidth, animation.Frame0Y
	screen.DrawImage(animation.SpriteSheet.SubImage(image.Rect(sx, sy, sx+animation.FrameWidth, sy+animation.FrameHeight)).(*ebiten.Image), op)
}
