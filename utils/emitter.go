package utils

type Event int

const (
	SWITCH_TO_ENEMY Event = iota
	SWITCH_TO_PLAYER
	SWITCH_TO_GAMEPLAY
)

type Emitter interface {
	Emit(event Event)
}
