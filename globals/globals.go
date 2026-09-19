package globals

import "tonic-quest/entities"

type Screen struct {
	Width  int
	Height int
}

type Globals struct {
	Enemy  *entities.Enemy
	Ally   *entities.Ally
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

func (g *Globals) SetEnemy(enemy *entities.Enemy) {
	g.Enemy = enemy
}

func (g *Globals) GetEnemy() *entities.Enemy {
	return g.Enemy
}

func (g *Globals) SetAlly(ally *entities.Ally) {
	g.Ally = ally
}

func (g *Globals) GetAlly() *entities.Ally {
	return g.Ally
}
