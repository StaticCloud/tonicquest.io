package main

import (
	"fmt"
	"log"
	"tonic-quest/enemy"
	"tonic-quest/keys"
	"tonic-quest/player"
	"tonic-quest/state"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game struct {
	SoundPlayer *keys.Player
	Enemy       *enemy.Enemy
	Player      *player.Player
	GameState   *state.Manager
}

var keyBuffer = []ebiten.Key{}

var flatKeys map[ebiten.Key]string = map[ebiten.Key]string{
	ebiten.KeyZ: "C",
	ebiten.KeyX: "D",
	ebiten.KeyC: "E",
	ebiten.KeyV: "F",
	ebiten.KeyB: "G",
	ebiten.KeyN: "A",
	ebiten.KeyM: "B",
}

func (g *Game) Update() error {
	if g.GameState.IsEnemyState() {
		g.Enemy.Attack()
	}

	if g.GameState.IsPlayerState() {
		move := g.Player.GetMove()

		if move != "" {
			g.Player.AttackList = append(g.Player.AttackList, move)

			if g.Player.AttackList[len(g.Player.AttackList)-1] != g.Enemy.AttackList[len(g.Player.AttackList)-1] {
				fmt.Println("Incorrect, the correct keys were: %s", g.Enemy.AttackList)
				g.Enemy.AttackList = []string{}
				g.Player.AttackList = []string{}
				g.GameState.SwitchToEnemyState()
			} else if len(g.Player.AttackList) == len(g.Enemy.AttackList) {
				fmt.Println("Correct!")
				g.Enemy.AttackList = []string{}
				g.Player.AttackList = []string{}
				g.GameState.SwitchToEnemyState()
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	context := audio.NewContext(48000)
	manager := state.InitStateManager()

	soundPlayer, err := keys.InitPlayer(context)
	if err != nil {
		panic(err)
	}

	enemy := enemy.InitEnemy([]string{"C", "D", "E", "F", "G", "A", "B"}, soundPlayer, manager)
	player := player.InitPlayer(soundPlayer, flatKeys, manager)

	game := &Game{
		SoundPlayer: soundPlayer,
		Enemy:       enemy,
		Player:      player,
		GameState:   manager,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Tonic Quest")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
