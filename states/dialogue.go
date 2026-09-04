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
	Orchestrator  *orchestrator.Orchestrator
}

func InitDialogueState(orchestrator *orchestrator.Orchestrator) *Dialogue {
	return &Dialogue{
		DialogueIndex: 0,
		CharIndex:     0,
		Column:        0,
		Dialogue:      "",
		Orchestrator:  orchestrator,
	}
}

func (s *Dialogue) Run() (StateName, error) {

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
		options.GeoM.Translate(0, 100)
		options.LineSpacing = 11 * 1.5

		current := ally.Dialogue[s.DialogueIndex]
		dialogue := current.Text
		speaker := current.Speaker

		if speaker == 0 {
			drawOptions := &ebiten.DrawImageOptions{}

			drawOptions.GeoM.Scale(0.6, 0.6)
			drawOptions.GeoM.Translate(0, 0)

			speakerOpts := &text.DrawOptions{}
			speakerOpts.GeoM.Translate(0, 80)

			text.Draw(screen, ally.Name, s.Orchestrator.Graphics.Font, speakerOpts)

			screen.DrawImage(
				s.Orchestrator.Graphics.Sprites[ally.Sprite],
				drawOptions,
			)
		}

		if speaker == 1 {
			drawOptions := &ebiten.DrawImageOptions{}

			drawOptions.GeoM.Scale(0.05, 0.05)
			drawOptions.GeoM.Translate(26, 0)

			speakerOpts := &text.DrawOptions{}
			speakerOpts.GeoM.Translate(0, 80)

			text.Draw(screen, "Carol", s.Orchestrator.Graphics.Font, speakerOpts)

			screen.DrawImage(
				s.Orchestrator.Graphics.Sprites["carol"],
				drawOptions,
			)
		}

		if s.CharIndex < len(dialogue) {
			s.Dialogue += string(dialogue[s.CharIndex])
			if s.Column > 35 && string(dialogue[s.CharIndex]) == " " {
				s.Dialogue += "\n"
				s.Column = 0
			} else {
				s.Column += 1
			}
			s.CharIndex += 1
		} else {
			keyBuffer := []ebiten.Key{}
			keyBuffer = inpututil.AppendPressedKeys(keyBuffer)

			if len(keyBuffer) > 0 {
				if keyBuffer[0] == ebiten.KeyN {
					s.DialogueIndex += 1
					s.reset()
				}
			}
		}

		text.Draw(screen, s.Dialogue, s.Orchestrator.Graphics.Font, options)
	}
}

func (s *Dialogue) reset() {
	s.CharIndex = 0
	s.Column = 0
	s.Dialogue = ""
}
