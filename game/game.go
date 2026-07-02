package game

import (
	"fmt"
	"tonic-quest/states"
	"tonic-quest/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type State int

const (
	SPLASH State = iota
	GAMEPLAY
	PLAYER_TURN
	ENEMY_TURN
	TERMINAL
)

type Manager struct {
	CurrentState State
	States       map[State]states.State
	Context      *utils.Context
}

func InitGameManager(context *utils.Context) *Manager {
	manager := &Manager{
		CurrentState: SPLASH,
		Context:      context,
		States: map[State]states.State{
			SPLASH:   states.InitSplashState(context),
			GAMEPLAY: states.InitGameplayState(context),
		},
	}

	manager.States[SPLASH].SetEmitter(manager)
	manager.States[GAMEPLAY].SetEmitter(manager)

	return manager
}

func (m *Manager) Run() error {
	err := m.States[m.CurrentState].Run()
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) Draw(screen *ebiten.Image) {
	m.States[m.CurrentState].Draw(screen)
}

func (m *Manager) Emit(event utils.Event) {
	switch {
	case event == utils.SWITCH_TO_GAMEPLAY:
		m.SwitchToGameplayState()
	}
}

func (m *Manager) LogGame() {
	fmt.Println("-----------------")
	fmt.Printf("Player Health: %d \n", m.Context.Player.Health)
	fmt.Printf("Enemy Health: %d \n", m.Context.Enemy.Health)
	fmt.Println("-----------------")
}

func (m *Manager) PlayKey(key string) {
	m.Context.Sound.Keys[key].Rewind()
	m.Context.Sound.Keys[key].Play()
}

func (m *Manager) NewTurn(newState func()) {
	m.LogGame()

	if m.Context.Enemy.Health <= 0 || m.Context.Player.Health <= 0 {
		m.SwitchToTerminalState()
	} else {
		m.Context.Enemy.AttackList = []string{}
		m.Context.Player.AttackList = []string{}

		newState()
	}
}

func (m *Manager) SwitchToGameplayState() {
	m.CurrentState = GAMEPLAY
}

func (m *Manager) SwitchToEnemyState() {
	m.CurrentState = ENEMY_TURN
}

func (m *Manager) SwitchToPlayerState() {
	m.CurrentState = PLAYER_TURN
}

func (m *Manager) SwitchToTerminalState() {
	m.CurrentState = TERMINAL
}

func (m *Manager) IsEnemyState() bool {
	return m.CurrentState == ENEMY_TURN
}

func (m *Manager) IsPlayerState() bool {
	return m.CurrentState == PLAYER_TURN
}

func (m *Manager) IsTerminalState() bool {
	return m.CurrentState == TERMINAL
}
