package overworld

type Enemy struct {
	X int
	Y int
}

func InitEnemy(x, y int) *Enemy {
	return &Enemy{
		X: x,
		Y: y,
	}
}
