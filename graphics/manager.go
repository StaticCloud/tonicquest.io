package graphics

import (
	"bytes"
	"image"
	_ "image/png"
	"io"
	"log"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
}

var sourceSprites map[string]string = map[string]string{
	"carol":               "./graphics/sprites/carol_sprite.png",
	"health":              "./graphics/sprites/combat/health.png",
	"health_border":       "./graphics/sprites/combat/health-border.png",
	"health_enemy_border": "./graphics/sprites/combat/health-enemy-border.png",
	"squelch":             "./graphics/sprites/squelch.png",
	"culo":                "./graphics/sprites/carol.png",
	"wall":                "./graphics/sprites/wall.png",
	"bunkus":              "./graphics/sprites/bunkus.png",
	"michael_cera":        "./graphics/sprites/michael_cera.png",
	"typewriter":          "./graphics/sprites/typewriter.png",
	"note":                "./graphics/sprites/note.png",
	"note_correct":        "./graphics/sprites/note_correct.png",
	"swipe":               "./graphics/sprites/swipe.png",
	"background":          "./graphics/sprites/combat/background.png",
	"culo_overworld":      "./graphics/sprites/overworld/culo.png",
	"dialogue_background": "./graphics/sprites/dialogue/text_background.png",
	"dialogue_backdrop":   "./graphics/sprites/dialogue/backdrop.png",
	"potion":              "./graphics/sprites/combat/potion.png",
	"shield":              "./graphics/sprites/combat/shield.png",
	"quantity_dot":        "./graphics/sprites/combat/literally_just_a_white_pixel.png",
}

var animatedSprites map[string]*Animation = map[string]*Animation{
	"carol_overworld":      InitAnimation("./graphics/sprites/bunkus.png", 20, 20, 5),
	"carol_idle":           InitAnimation("./graphics/sprites/combat/carol_idle.png", 120, 120, 2),
	"culo_idle":            InitAnimation("./graphics/sprites/combat/culo_idle.png", 120, 120, 2),
	"carol_neutral":        InitAnimation("./graphics/sprites/dialogue/carol_neutral.png", 40, 40, 2),
	"michael_cera_neutral": InitAnimation("./graphics/sprites/dialogue/michael_cera.png", 40, 40, 2),
}

var backgrounds map[string]*ebiten.Image = map[string]*ebiten.Image{}
var sprites map[string]*ebiten.Image = map[string]*ebiten.Image{}

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

	alagard, err := os.Open("./graphics/fonts/alagard.ttf")
	if err != nil {
		return nil, err
	}

	defer alagard.Close()

	byteValue, byteError := io.ReadAll(alagard)
	if byteError != nil {
		return nil, byteError
	}

	source, err := text.NewGoTextFaceSource(bytes.NewReader(byteValue))

	if err != nil {
		log.Fatal(err)
	}

	face := &text.GoTextFace{
		Source:    source,
		Size:      16,
		Direction: text.DirectionLeftToRight,
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

func (m *Manager) DrawAnimation(sprite string, x, y int, screen *ebiten.Image, speed float64) {
	animation := animatedSprites[sprite]
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	tempo := int(math.Floor(float64(m.Ticker) * speed))
	i := (tempo / 5) % animation.FrameCount
	sx, sy := animation.Frame0X+i*animation.FrameWidth, animation.Frame0Y
	screen.DrawImage(animation.SpriteSheet.SubImage(image.Rect(sx, sy, sx+animation.FrameWidth, sy+animation.FrameHeight)).(*ebiten.Image), op)
}
