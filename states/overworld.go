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
	s.Player.ReadInput(s.getValidInputs())

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
	s.renderOverworld(screen)
	playerOptions := &ebiten.DrawImageOptions{}

	playerOptions.GeoM.Translate(float64(s.Player.X), float64(s.Player.Y))

	image := ebiten.NewImage(20, 20)
	image.Fill(color.White)

	screen.DrawImage(
		s.Orchestrator.Graphics.Sprites["bunkus"],
		playerOptions,
	)

	for _, enemy := range s.World.Locations[0].Enemies {
		enemyOptions := &ebiten.DrawImageOptions{}

		enemyOptions.GeoM.Scale(0.2, 0.2)
		enemyOptions.GeoM.Translate(float64(enemy.Coordinates.X)*20, float64(enemy.Coordinates.Y)*20)

		screen.DrawImage(
			s.Orchestrator.Graphics.Sprites[enemy.Sprite],
			enemyOptions,
		)
	}
}

func (s *Overworld) getValidInputs() map[ebiten.Key]bool {
	x := s.Player.X / 20
	y := s.Player.Y / 20

	allowed := make(map[ebiten.Key]bool)
	// ensure coordinate is within allowable bounds
	if y-1 >= 0 {
		if s.World.Locations[0].Map[y-1][x] == "X" {
			allowed[ebiten.KeyW] = false
		} else {
			allowed[ebiten.KeyW] = true
		}
	}

	if x-1 >= 0 {
		if s.World.Locations[0].Map[y][x-1] == "X" {
			allowed[ebiten.KeyA] = false
		} else {
			allowed[ebiten.KeyA] = true
		}
	}

	if y+1 < 8 {
		if s.World.Locations[0].Map[y+1][x] == "X" {
			allowed[ebiten.KeyS] = false
		} else {
			allowed[ebiten.KeyS] = true
		}
	}

	if x+1 < 12 {
		if s.World.Locations[0].Map[y][x+1] == "X" {
			allowed[ebiten.KeyD] = false
		} else {
			allowed[ebiten.KeyD] = true
		}
	}

	return allowed
}

func (s *Overworld) renderOverworld(screen *ebiten.Image) {

	for i, _ := range s.World.Locations[0].Map {
		for j, _ := range s.World.Locations[0].Map[i] {
			block := s.World.Locations[0].Map[i][j]
			if block == "X" {
				blockade := &ebiten.DrawImageOptions{}

				blockade.GeoM.Scale(1, 1)
				blockade.GeoM.Translate(float64(j)*20, float64(i)*20)

				screen.DrawImage(
					s.Orchestrator.Graphics.Sprites["twee"],
					blockade,
				)
			}
		}
	}
}
