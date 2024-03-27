package speech

import "github.com/RacoonMediaServer/rms-packages/pkg/worker"

type Workers interface {
	Do(t worker.Task) worker.Receipt
	DoneChannel() <-chan string
}
