package keys

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var keys []string = []string{
	"C", "D", "E", "F", "G", "A", "B",
}

type Player struct {
	Context *audio.Context
	Keys    map[string]*audio.Player
}

func LoadFile(name string, context *audio.Context) (*audio.Player, error) {
	file, fileErr := os.Open(name)
	if fileErr != nil {
		return nil, fileErr
	}

	decodedWav, decodedWavErr := wav.DecodeF32(file)
	if decodedWavErr != nil {
		return nil, decodedWavErr
	}

	return context.NewPlayerF32(decodedWav)

}

func InitPlayer(context *audio.Context) (*Player, error) {
	keyDict := map[string]*audio.Player{}

	for _, key := range keys {
		keyDict[key], _ = LoadFile(fmt.Sprintf("./keys/%s.wav", key), context)
	}

	return &Player{
		Context: context,
		Keys:    keyDict,
	}, nil
}
