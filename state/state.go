package state

type State int

const (
	PLAYER_TURN State = 0
	ENEMY_TURN  State = 1
)

type Manager struct {
	CurrentState State
}

func InitStateManager() *Manager {
	return &Manager{
		CurrentState: ENEMY_TURN,
	}
}

func (m *Manager) SwitchToEnemyState() {
	m.CurrentState = ENEMY_TURN
}

func (m *Manager) SwitchToPlayerState() {
	m.CurrentState = PLAYER_TURN
}

func (m *Manager) IsEnemyState() bool {
	return m.CurrentState == ENEMY_TURN
}

func (m *Manager) IsPlayerState() bool {
	return m.CurrentState == PLAYER_TURN
}
