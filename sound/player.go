package sound

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

func InitPlayer() (*Player, error) {
	context := audio.NewContext(48000)
	keyDict := map[string]*audio.Player{}

	for _, key := range keys {
		keyDict[key], _ = loadFile(fmt.Sprintf("./sound/%s.wav", key), context)
	}

	return &Player{
		Context: context,
		Keys:    keyDict,
	}, nil
}

func loadFile(name string, context *audio.Context) (*audio.Player, error) {
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
