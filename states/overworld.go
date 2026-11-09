package states

import (
	"fmt"
	"math"
	"strings"
	data "tonic-quest/data/overworld"
	"tonic-quest/entities/combat"
	"tonic-quest/entities/overworld"
	"tonic-quest/entities/utils"
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Event int

const (
	DEFAULT Event = iota
	DIALOGUE
	TRANSITION_OUT
)

type Overworld struct {
	Orchestrator *orchestrator.Orchestrator
	Player       *overworld.Player
	Enemies      []*overworld.Enemy
	World        *data.Data
	Ticker       *utils.Timer
	Event        Event
	Swipe        *overworld.Swipe
	Events       map[Event]func()
	Dialogue     *overworld.Dialogue
}

type Reset struct {
	Overworld *Overworld
}

func InitOverworldState(orchestrator *orchestrator.Orchestrator) *Overworld {
	locations, err := data.LoadWorld()

	if err != nil {
		fmt.Println(err)
	}

	overworld := &Overworld{
		Orchestrator: orchestrator,
		Player:       overworld.InitPlayer(),
		World:        locations,
		Ticker:       utils.InitTimer(),
		Event:        DEFAULT,
		Swipe:        overworld.InitSwipe(),
		Dialogue:     overworld.InitDialogue(),
	}

	overworld.Events = map[Event]func(){
		DEFAULT:        overworld.Overworld,
		TRANSITION_OUT: overworld.TransitionOut,
		DIALOGUE:       overworld.DialogueActive,
	}

	return overworld
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
	s.Events[s.Event]()

	s.Orchestrator.Graphics.Update()
	if s.Event == DIALOGUE {
		s.Dialogue.Update()
	}

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
			s.Orchestrator.Globals.SetEnemy(combat.InitEnemy(
				enemy.ID,
				enemy.Health,
				enemy.Name,
				enemy.Sprite,
				enemy.Moves,
			))
			s.Orchestrator.Graphics.Ticker = 0
			s.Event = TRANSITION_OUT

			if s.Ticker.Tick == 20 {
				return LOAD_COMBAT, nil
			}
		}
	}

	for _, ally := range s.World.Locations[0].Allies {
		if ally.Coordinates.X*20 == s.Player.X && ally.Coordinates.Y*20 == s.Player.Y {
			keyBuffer := []ebiten.Key{}
			keyBuffer = inpututil.AppendPressedKeys(keyBuffer)

			if len(keyBuffer) > 0 {
				if keyBuffer[0] == ebiten.KeySpace {
					s.Orchestrator.Globals.SetAlly(&overworld.Ally{
						Name:     ally.Name,
						Sprite:   ally.Sprite,
						Dialogue: ally.Dialogue,
					})
					s.SetEvent(DIALOGUE)
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
	// playerOptions := &ebiten.DrawImageOptions{}

	// playerOptions.GeoM.Translate(float64(s.Player.X), float64(s.Player.Y))

	// screen.DrawImage(
	// 	s.Orchestrator.Graphics.Sprites["bunkus"],
	// 	playerOptions,
	// )

	s.Orchestrator.Graphics.DrawAnimation("carol_overworld", s.Player.X, s.Player.Y, screen, 1.0)

	for _, enemy := range s.World.Locations[0].Enemies {
		if enemy.Defeated == false {
			enemyOptions := &ebiten.DrawImageOptions{}

			enemyOptions.GeoM.Translate(float64(enemy.Coordinates.X)*20, float64(enemy.Coordinates.Y)*20)

			screen.DrawImage(
				s.Orchestrator.Graphics.Sprites[fmt.Sprintf("%s_overworld", enemy.Sprite)],
				enemyOptions,
			)
		}
	}

	for _, ally := range s.World.Locations[0].Allies {
		allyOptions := &ebiten.DrawImageOptions{}

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

	swipeOptions := &ebiten.DrawImageOptions{}
	swipeOptions.GeoM.Translate(float64(s.Swipe.X), 0)

	screen.DrawImage(
		s.Orchestrator.Graphics.Sprites["swipe"],
		swipeOptions,
	)

	if s.Event == DIALOGUE {
		dialogueBackdrop := &ebiten.DrawImageOptions{}

		dialogueBackdrop.GeoM.Translate(float64(s.Dialogue.Coords.X), 0)

		screen.DrawImage(
			s.Orchestrator.Graphics.Sprites["dialogue_backdrop"],
			dialogueBackdrop,
		)

		dialogueBackground := &ebiten.DrawImageOptions{}

		dialogueBackground.GeoM.Translate(float64(s.Dialogue.Coords.X), 0)

		screen.DrawImage(
			s.Orchestrator.Graphics.Sprites["dialogue_background"],
			dialogueBackground,
		)

		ally := s.Orchestrator.Globals.Ally

		current := ally.Dialogue[s.Dialogue.Index]

		options := &text.DrawOptions{}
		options.GeoM.Translate(float64(s.Dialogue.Coords.X)+10, 55)
		options.LineSpacing = 11 * 1.5

		sprite := s.generateSprite(current.Speaker.Name, current.Speaker.Emotion)

		s.Orchestrator.Graphics.DrawAnimation(sprite, s.Dialogue.Coords.X+10, 10, screen, float64(s.Dialogue.Speed))
		text.Draw(screen, s.Dialogue.Dialogue, s.Orchestrator.Graphics.Font, options)
	}
}

func (s *Overworld) generateSprite(speaker, emotion string) string {
	return fmt.Sprintf("%s_%s", strings.ToLower(speaker), emotion)
}

func (s *Overworld) SetEvent(event Event) {
	s.Event = event
}

func (s *Overworld) Overworld() {
	s.Player.ReadInput(s.getValidInputs())
}

func (s *Overworld) TransitionOut() {
	s.Ticker.Update()
	s.Swipe.Update()
	s.Swipe.Start()
}

func (s *Overworld) DialogueActive() {
	ally := s.Orchestrator.Globals.Ally

	if s.Dialogue.Index+1 == len(ally.Dialogue) && s.Dialogue.IsInactive() {
		s.Event = DEFAULT
	}

	if s.Dialogue.Index == 0 && s.Dialogue.IsInactive() {
		s.Dialogue.InvokeEntrance()
	}

	if s.Dialogue.Index < len(ally.Dialogue) {
		current := ally.Dialogue[s.Dialogue.Index]
		dialogue := current.Text

		if current.Pause == true {
			s.Dialogue.Speed = 0
		}

		if s.Dialogue.Char < len(dialogue) {
			s.Dialogue.Dialogue += string(dialogue[s.Dialogue.Char])
			if s.Dialogue.Column >= 20 && string(dialogue[s.Dialogue.Char]) == " " {
				s.Dialogue.Dialogue += "\n"
				s.Dialogue.Column = 0
			} else {
				s.Dialogue.Column += 1
			}
			s.Dialogue.Char += 1
		} else {
			s.Dialogue.Speed = 0
			keyBuffer := []ebiten.Key{}
			keyBuffer = inpututil.AppendPressedKeys(keyBuffer)

			if len(keyBuffer) > 0 {
				if keyBuffer[0] == ebiten.KeyN {
					if s.Dialogue.Index+1 != len(ally.Dialogue) {
						s.Dialogue.Next()
					} else {
						s.Dialogue.Reset(s.Default)
					}
				}
			}
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
					s.Orchestrator.Graphics.Sprites["wall"],
					blockade,
				)
			}
		}
	}
}

func (s *Overworld) renderText(screen *ebiten.Image, displayText string) {
	speakerOpts := &text.DrawOptions{}

	w, _ := text.Measure(displayText, s.Orchestrator.Graphics.Font, 0)
	screenW := screen.Bounds().Dx()

	posX := math.Floor(float64(screenW-int(math.Floor(w))) / 2.0)

	speakerOpts.GeoM.Translate(posX, 140)
	text.Draw(screen, displayText, s.Orchestrator.Graphics.Font, speakerOpts)
}

func (s *Overworld) Default() {
	s.Event = DEFAULT
}
