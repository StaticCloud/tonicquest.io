package state

import (
	"fmt"
	"tonic-quest/entities"
	"tonic-quest/keys"
)

type State int

const (
	PLAYER_TURN State = 0
	ENEMY_TURN  State = 1
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

func (m *Manager) Run() {
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
				fmt.Println("Incorrect, the correct keys were: ", m.Enemy.AttackList)
				m.NewTurn(m.SwitchToEnemyState)
			} else if len(m.Player.AttackList) == len(m.Enemy.AttackList) {
				fmt.Println("Correct!")
				m.NewTurn(m.SwitchToEnemyState)
			}
		}
	}
}

func (m *Manager) PlayKey(key string) {
	m.SoundManager.Keys[key].Rewind()
	m.SoundManager.Keys[key].Play()
}

func (m *Manager) NewTurn(newState func()) {
	m.Enemy.AttackList = []string{}
	m.Player.AttackList = []string{}

	newState()
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
