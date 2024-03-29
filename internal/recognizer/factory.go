package recognizer

import (
	"context"
	"github.com/RacoonMediaServer/rms-speech/internal/recognizer/whisper"
	"golang.org/x/text/language"
)

type Recognizer interface {
	Transcribe(ctx context.Context, inputFile string, lang language.Tag) (string, error)
}

type Type int

const (
	WhisperExecutable Type = iota
	WhisperAPI
)

func New(t Type, setting string) Recognizer {
	switch t {
	case WhisperExecutable:
		return whisper.NewExecutableRecognizer(setting)
	default:
		panic("not implemented")
	}
}
