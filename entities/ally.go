package entities

import data "tonic-quest/data/overworld"

type Ally struct {
	Name     string
	Sprite   string
	Dialogue []data.Dialogue
}

func InitAlly(name, sprite string, dialogue []data.Dialogue) *Ally {
	return &Ally{
		Name:     name,
		Sprite:   sprite,
		Dialogue: dialogue,
	}
}
