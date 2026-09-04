package globals

import "tonic-quest/entities"

type Globals struct {
	Enemy *entities.Enemy
	Ally  *entities.Ally
}

func InitGlobals() (*Globals, error) {
	return &Globals{}, nil
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
