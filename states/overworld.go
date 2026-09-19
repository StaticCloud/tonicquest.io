package states

import (
	"fmt"
	"image/color"
	data "tonic-quest/data/overworld"
	"tonic-quest/entities"
	"tonic-quest/entities/overworld"
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Overworld struct {
	Orchestrator *orchestrator.Orchestrator
	Player       *overworld.Player
	Enemies      []*overworld.Enemy
	World        *data.Data
}

type Reset struct {
	Overworld *Overworld
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

func RefreshOverworldState(overworld *Overworld) *Reset {
	return &Reset{
		Overworld: overworld,
	}
}

func (r *Reset) Run() (StateName, error) {
	r.Overworld.Player.Reset()

	return OVERWORLD, nil
}

func (r *Reset) Draw(screen *ebiten.Image) {}

func (s *Overworld) Run() (StateName, error) {
	s.Player.ReadInput(s.getValidInputs())

	if s.Orchestrator.Globals.Enemy != nil {
		if s.Orchestrator.Globals.Enemy.Defeated == true {
			s.World.Locations[0].Enemies[s.Orchestrator.Globals.Enemy.ID-1].Defeated = true
			data.SaveWorld(s.World)
			s.Orchestrator.Globals.Enemy = nil
		}

		s.Orchestrator.Globals.Enemy = nil
	}

	for _, enemy := range s.World.Locations[0].Enemies {
		if (enemy.Coordinates.X*20 == s.Player.X && enemy.Coordinates.Y*20 == s.Player.Y) && enemy.Defeated == false {
			s.Orchestrator.Globals.SetEnemy(&entities.Enemy{
				ID:     enemy.ID,
				Name:   enemy.Name,
				Sprite: enemy.Sprite,
				Health: 100,
				Moves:  enemy.Moves,
			})
			return GAMEPLAY, nil
		}
	}

	for _, ally := range s.World.Locations[0].Allies {
		if ally.Coordinates.X*20 == s.Player.X && ally.Coordinates.Y*20 == s.Player.Y {
			keyBuffer := []ebiten.Key{}
			keyBuffer = inpututil.AppendPressedKeys(keyBuffer)

			if len(keyBuffer) > 0 {
				if keyBuffer[0] == ebiten.KeySpace {
					s.Orchestrator.Globals.SetAlly(&entities.Ally{
						Name:     ally.Name,
						Sprite:   ally.Sprite,
						Dialogue: ally.Dialogue,
					})
					return DIALOGUE, nil
				}
			}
		}
	}

	for _, savepoint := range s.World.Locations[0].Savepoints {
		if savepoint.Coordinates.X*20 == s.Player.X && savepoint.Coordinates.Y*20 == s.Player.Y {
			keyBuffer := []ebiten.Key{}
			keyBuffer = inpututil.AppendPressedKeys(keyBuffer)

			if len(keyBuffer) > 0 {
				if keyBuffer[0] == ebiten.KeySpace {
					data.SaveWorld(s.World)
				}
			}
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
		if enemy.Defeated == false {
			enemyOptions := &ebiten.DrawImageOptions{}

			enemyOptions.GeoM.Scale(0.2, 0.2)
			enemyOptions.GeoM.Translate(float64(enemy.Coordinates.X)*20, float64(enemy.Coordinates.Y)*20)

			screen.DrawImage(
				s.Orchestrator.Graphics.Sprites[enemy.Sprite],
				enemyOptions,
			)
		}
	}

	for _, ally := range s.World.Locations[0].Allies {
		allyOptions := &ebiten.DrawImageOptions{}

		allyOptions.GeoM.Scale(0.2, 0.2)
		allyOptions.GeoM.Translate(float64(ally.Coordinates.X)*20, float64(ally.Coordinates.Y)*20)

		screen.DrawImage(
			s.Orchestrator.Graphics.Sprites[ally.Sprite],
			allyOptions,
		)

		if ally.Coordinates.X*20 == s.Player.X && ally.Coordinates.Y*20 == s.Player.Y {
			s.renderText(screen, fmt.Sprintf("Talk to %s", ally.Name))
		}
	}

	for _, savepoint := range s.World.Locations[0].Savepoints {
		allyOptions := &ebiten.DrawImageOptions{}

		allyOptions.GeoM.Translate(float64(savepoint.Coordinates.X)*20, float64(savepoint.Coordinates.Y)*20)

		screen.DrawImage(
			s.Orchestrator.Graphics.Sprites["typewriter"],
			allyOptions,
		)

		if savepoint.Coordinates.X*20 == s.Player.X && savepoint.Coordinates.Y*20 == s.Player.Y {
			s.renderText(screen, "Save Game")
		}
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

func (s *Overworld) renderText(screen *ebiten.Image, displayText string) {
	speakerOpts := &text.DrawOptions{}
	speakerOpts.LayoutOptions.PrimaryAlign = text.AlignCenter
	speakerOpts.GeoM.Translate(120, 140)

	text.Draw(screen, displayText, s.Orchestrator.Graphics.Font, speakerOpts)
}
