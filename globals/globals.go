package globals

import "tonic-quest/entities"

type Globals struct {
	Enemy *entities.Enemy
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
