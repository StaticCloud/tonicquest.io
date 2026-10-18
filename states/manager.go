package states

import (
	"tonic-quest/orchestrator"

	"github.com/hajimehoshi/ebiten/v2"
)

type StateName int
type SubState int

const (
	NONE StateName = iota
	SPLASH
	COMBAT
	DIALOGUE
	PLAYER_TURN
	ENEMY_TURN
	TERMINAL
	OVERWORLD
	OVERWORLD_LOSE
	LOAD_COMBAT
)

const (
	STATE SubState = iota
	TRANSITION_OUT
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
	manager := &Manager{
		States: map[StateName]State{
			OVERWORLD: InitOverworldState(orchestrator),
			SPLASH:    InitSplashState(orchestrator),
			COMBAT:    InitGameplayState(orchestrator),
			DIALOGUE:  InitDialogueState(orchestrator),
		},
		State: SPLASH,
	}

	manager.States[OVERWORLD_LOSE] = RefreshOverworldState(manager.States[OVERWORLD].(*Overworld))
	manager.States[LOAD_COMBAT] = RefreshCombatState(manager.States[COMBAT].(*Gameplay))

	return manager
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
