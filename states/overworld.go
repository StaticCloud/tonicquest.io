package states

import (
	"fmt"
	"image/color"
	data "tonic-quest/data/overworld"
	"tonic-quest/entities"
	"tonic-quest/entities/overworld"
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
)

type Overworld struct {
	Orchestrator *orchestrator.Orchestrator
	Player       *overworld.Player
	Enemies      []*overworld.Enemy
	World        *data.Data
}

func InitOverworldState(orchestrator *orchestrator.Orchestrator) *Overworld {
	locations, err := data.LoadWorld()
	if err != nil {
		fmt.Println(err)
	}
	return &Overworld{
		Orchestrator: orchestrator,
		Player:       overworld.InitPlayer(),
		World:        locations,
	}
}

func (s *Overworld) Run() (StateName, error) {
	s.Player.ReadInput()

	for _, enemy := range s.World.Locations[0].Enemies {
		if enemy.Coordinates.X*20 == s.Player.X && enemy.Coordinates.Y*20 == s.Player.Y {
			s.Orchestrator.Globals.SetEnemy(&entities.Enemy{
				Name:   enemy.Name,
				Sprite: enemy.Sprite,
				Health: 100,
				Moves:  []string{"C", "D", "E"},
			})
			return GAMEPLAY, nil
		}
	}

	return OVERWORLD, nil
}

func (s *Overworld) Draw(screen *ebiten.Image) {
	playerOptions := &ebiten.DrawImageOptions{}

	playerOptions.GeoM.Translate(float64(s.Player.X), float64(s.Player.Y))

	image := ebiten.NewImage(20, 20)
	image.Fill(color.White)

	screen.DrawImage(
		image,
		playerOptions,
	)

	for _, enemy := range s.World.Locations[0].Enemies {
		enemyOptions := &ebiten.DrawImageOptions{}

		enemyOptions.GeoM.Translate(float64(enemy.Coordinates.X)*20, float64(enemy.Coordinates.Y)*20)

		image := ebiten.NewImage(20, 20)
		image.Fill(color.RGBA{255, 0, 0, 0})

		screen.DrawImage(
			image,
			enemyOptions,
		)
	}
}
