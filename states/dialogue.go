package states

import (
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Dialogue struct {
	DialogueIndex int
	CharIndex     int
	Column        int
	Dialogue      string
	Speed         float64
	Orchestrator  *orchestrator.Orchestrator
}

func InitDialogueState(orchestrator *orchestrator.Orchestrator) *Dialogue {
	return &Dialogue{
		DialogueIndex: 0,
		CharIndex:     0,
		Column:        0,
		Dialogue:      "",
		Speed:         1,
		Orchestrator:  orchestrator,
	}
}

func (s *Dialogue) Run() (StateName, error) {
	s.Orchestrator.Graphics.Update()
	if len(s.Orchestrator.Globals.Ally.Dialogue) <= s.DialogueIndex {
		s.DialogueIndex = 0
		return OVERWORLD, nil
	}

	return DIALOGUE, nil
}

func (s *Dialogue) Draw(screen *ebiten.Image) {
	ally := s.Orchestrator.Globals.Ally

	if s.DialogueIndex < len(ally.Dialogue) {
		options := &text.DrawOptions{}
		options.GeoM.Translate(0, 80)
		options.LineSpacing = 11 * 1.5

		current := ally.Dialogue[s.DialogueIndex]
		dialogue := current.Text
		speaker := current.Speaker

		if current.Pause == true {
			s.Speed = 0
		}

		s.Orchestrator.Graphics.DrawAnimation(speaker.Sprite, 0, 0, screen, float64(s.Speed))

		if s.CharIndex < len(dialogue) {
			s.Dialogue += string(dialogue[s.CharIndex])
			if s.Column > 25 && string(dialogue[s.CharIndex]) == " " {
				s.Dialogue += "\n"
				s.Column = 0
			} else {
				s.Column += 1
			}
			s.CharIndex += 1
		} else {
			s.Speed = 0
			keyBuffer := []ebiten.Key{}
			keyBuffer = inpututil.AppendPressedKeys(keyBuffer)

			if len(keyBuffer) > 0 {
				if keyBuffer[0] == ebiten.KeyN {
					s.DialogueIndex += 1
					s.Reset()
				}
			}
		}

		text.Draw(screen, s.Dialogue, s.Orchestrator.Graphics.Font, options)
	}
}

func (s *Dialogue) Reset() {
	s.CharIndex = 0
	s.Column = 0
	s.Dialogue = ""
	s.Speed = 1
}
