package state

import (
	"fmt"
	"tonic-quest/entities"
	"tonic-quest/keys"
)

type State int

const (
	PLAYER_TURN State = iota
	ENEMY_TURN
	TERMINAL
)

type Manager struct {
	CurrentState State
	SoundManager *keys.Player
	Player       *entities.Player
	Enemy        *entities.Enemy
}

func InitStateManager(soundManager *keys.Player, player *entities.Player, enemy *entities.Enemy) *Manager {
	return &Manager{
		CurrentState: ENEMY_TURN,
		SoundManager: soundManager,
		Player:       player,
		Enemy:        enemy,
	}
}

func (m *Manager) Run() error {
	if m.IsEnemyState() {
		for range 3 {
			move := m.Enemy.Attack()
			m.PlayKey(move)
		}

		m.SwitchToPlayerState()
	}

	if m.IsPlayerState() {
		move := m.Player.Attack()

		if move != "" {
			m.Player.AttackList = append(m.Player.AttackList, move)

			if m.Player.AttackList[len(m.Player.AttackList)-1] != m.Enemy.AttackList[len(m.Player.AttackList)-1] {
				m.Player.TakeDamage(10)
				m.NewTurn(m.SwitchToEnemyState)
			} else if len(m.Player.AttackList) == len(m.Enemy.AttackList) {
				m.Enemy.TakeDamage(10)
				m.NewTurn(m.SwitchToEnemyState)
			}
		}
	}

	if m.IsTerminalState() {
		return fmt.Errorf("Terminated")
	}

	return nil
}

func (m *Manager) LogGame() {
	fmt.Println("-----------------")
	fmt.Printf("Player Health: %d \n", m.Player.Health)
	fmt.Printf("Enemy Health: %d \n", m.Enemy.Health)
	fmt.Println("-----------------")
}

func (m *Manager) PlayKey(key string) {
	m.SoundManager.Keys[key].Rewind()
	m.SoundManager.Keys[key].Play()
}

func (m *Manager) NewTurn(newState func()) {
	m.LogGame()

	if m.Enemy.Health <= 0 || m.Player.Health <= 0 {
		m.SwitchToTerminalState()
	} else {
		m.Enemy.AttackList = []string{}
		m.Player.AttackList = []string{}

		newState()
	}
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
