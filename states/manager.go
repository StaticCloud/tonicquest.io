package states

import (
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
)

type StateName int

const (
	NONE StateName = iota
	SPLASH
	GAMEPLAY
	PLAYER_TURN
	ENEMY_TURN
	TERMINAL
)

type State interface {
	Run() (StateName, error)
	Draw(screen *ebiten.Image)
}

type Manager struct {
	State  StateName
	States map[StateName]State
}

func InitManager(orchestrator *orchestrator.Orchestrator) *Manager {
	return &Manager{
		States: map[StateName]State{
			SPLASH:   InitSplashState(orchestrator),
			GAMEPLAY: InitGameplayState(orchestrator),
		},
		State: SPLASH,
	}
}

func (s *Manager) Run() error {
	state, err := s.States[s.State].Run()

	if err != nil {
		return err
	}

	if state != NONE {
		s.SwitchState(state)
	}

	return nil
}

func (s *Manager) Draw(screen *ebiten.Image) {
	s.States[s.State].Draw(screen)
}

func (m *Manager) SwitchState(state StateName) {
	m.State = state
}

func (m *Manager) IsEnemyState() bool {
	return m.State == ENEMY_TURN
}

func (m *Manager) IsPlayerState() bool {
	return m.State == PLAYER_TURN
}

func (m *Manager) IsTerminalState() bool {
	return m.State == TERMINAL
}
