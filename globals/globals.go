package globals

import (
	"tonic-quest/entities/combat"
	"tonic-quest/entities/overworld"
)

type Screen struct {
	Width  int
	Height int
}

type Globals struct {
	Enemy  *combat.Enemy
	Ally   *overworld.Ally
	Screen Screen
}

func InitGlobals() (*Globals, error) {
	return &Globals{
		Screen: Screen{
			Width:  240,
			Height: 160,
		},
	}, nil
}

func (g *Globals) SetEnemy(enemy *combat.Enemy) {
	g.Enemy = enemy
}

func (g *Globals) GetEnemy() *combat.Enemy {
	return g.Enemy
}

func (g *Globals) SetAlly(ally *overworld.Ally) {
	g.Ally = ally
}

func (g *Globals) GetAlly() *overworld.Ally {
	return g.Ally
}
