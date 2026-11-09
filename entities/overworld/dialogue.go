package overworld

import (
	"math"
	"tonic-quest/entities/utils"
)

type Coords struct {
	X int
	Y int
}

type Event int

const (
	INACTIVE Event = iota
	ENTERING
	ACTIVE
)

type Dialogue struct {
	Index            int
	Char             int
	Column           int
	Dialogue         string
	Speed            float64
	AnimationManager utils.AnimationManager
	Coords           *Coords
	Event            Event
	EventDictionary  map[Event]func()
}

func InitDialogue() *Dialogue {

	dialogue := &Dialogue{
		Index:            0,
		Char:             0,
		Column:           0,
		Dialogue:         "",
		Speed:            1,
		AnimationManager: *utils.InitAnimationManager(),
		Coords: &Coords{
			X: 240,
			Y: 0,
		},
		Event: INACTIVE,
	}

	eventDictionary := map[Event]func(){
		INACTIVE: dialogue.Inactive,
		ACTIVE:   dialogue.Active,
		ENTERING: dialogue.Entrance,
	}

	dialogue.EventDictionary = eventDictionary

	return dialogue
}

func (d *Dialogue) Update() error {
	if d.AnimationManager.Pause == false {
		d.AnimationManager.Tick += 1
	}
	d.Animate()
	return nil
}

func (d *Dialogue) InvokeActive() {
	d.Event = ACTIVE
}

func (d *Dialogue) InvokeInactive() {
	d.Event = INACTIVE
}

func (d *Dialogue) InvokeEntrance() {
	d.Event = ENTERING
}

func (d *Dialogue) IsActive() bool {
	return d.Event == ACTIVE
}

func (d *Dialogue) IsInactive() bool {
	return d.Event == INACTIVE
}

func (d *Dialogue) IsEntering() bool {
	return d.Event == ENTERING
}

func (d *Dialogue) SetX(x int) {
	d.Coords.X = x
}

func (d *Dialogue) Animate() {
	d.EventDictionary[d.Event]()
}

func (d *Dialogue) Inactive() {
	d.SetX(240)
}

func (d *Dialogue) Active() {
	d.AnimationManager.Pause = true
	d.SetX(0)
}

func (d *Dialogue) Entrance() {
	d.AnimationManager.Pause = false
	d.AnimateEntrance()
}

func (d *Dialogue) AnimateEntrance() {
	x := int((d.AnimationManager.Calculate(math.Sin, 240, -0.05, 240, 1)))

	if x != 0 {
		d.SetX(x)
	} else {
		d.InvokeActive()
	}
}

func (d *Dialogue) Reset(effect func()) {
	d.Index = 0
	d.Char = 0
	d.Column = 0
	d.Dialogue = ""
	d.Speed = 1
	d.SetX(240)

	d.AnimationManager.Tick = 0
	d.InvokeInactive()
	effect()
}

func (d *Dialogue) Next() {
	d.Index += 1
	d.Char = 0
	d.Column = 0
	d.Dialogue = ""
	d.Speed = 1
}
