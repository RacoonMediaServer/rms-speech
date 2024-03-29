package speech

import (
	"context"
	"github.com/RacoonMediaServer/rms-packages/pkg/worker"
	"github.com/RacoonMediaServer/rms-speech/internal/recognizer"
	"golang.org/x/text/language"
	"os"
)

type job struct {
	receipt worker.Receipt
	t       *task
}

type task struct {
	id             string
	inputFile      string
	recognizedText string
	recognizer     recognizer.Recognizer
}

func (t *task) ID() string {
	return t.id
}

func (t *task) Do(ctx context.Context) error {
	defer t.clean()
	var err error

	// TODO: specify language to API
	t.recognizedText, err = t.recognizer.Transcribe(ctx, t.inputFile, language.Russian)
	return err
}

func (t *task) clean() {
	_ = os.Remove(t.inputFile)
}
