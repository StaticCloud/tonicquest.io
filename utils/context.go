package utils

import (
	"tonic-quest/entities"
	"tonic-quest/graphics"
	"tonic-quest/keys"
)

type Context struct {
	Player   *entities.Player
	Enemy    *entities.Enemy
	Sound    *keys.Player
	Graphics *graphics.BackgroundManager
}
