package speech

import (
	"context"
	"github.com/RacoonMediaServer/rms-packages/pkg/worker"
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
}

func (t *task) ID() string {
	return t.id
}

func (t *task) Do(ctx context.Context) error {
	defer t.clean()
	//TODO implement me
	panic("implement me")
}

func (t *task) clean() {
	_ = os.Remove(t.inputFile)
}
