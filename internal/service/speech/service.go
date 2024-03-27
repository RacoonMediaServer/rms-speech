package speech

import (
	"context"
	"errors"
	rms_speech "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-speech"
	"github.com/RacoonMediaServer/rms-packages/pkg/worker"
	"github.com/google/uuid"
	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"path/filepath"
	"sync"
)

type Service struct {
	workers Workers

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	l    logger.Logger
	mu   sync.Mutex
	jobs map[string]*job
}

func New(workers Workers) *Service {
	s := Service{
		workers: workers,
		jobs:    map[string]*job{},
		l:       logger.DefaultLogger.Fields(map[string]interface{}{"from": "speech"}),
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.processReadyJobs()
	}()

	return &s
}

func (s *Service) StartRecognition(ctx context.Context, request *rms_speech.StartRecognitionRequest, response *rms_speech.StartRecognitionResponse) error {
	id, err := uuid.NewUUID()
	if err != nil {
		s.l.Logf(logger.ErrorLevel, "Generate id failed: %s", err)
		return err
	}

	fileName := filepath.Join(os.TempDir(), id.String())
	if err = os.WriteFile(fileName, request.Data, 0666); err != nil {
		s.l.Logf(logger.ErrorLevel, "Save audio file failed: %s", err)
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	t := &task{
		id:        id.String(),
		inputFile: fileName,
	}
	j := &job{
		receipt: s.workers.Do(t),
		t:       t,
	}
	s.jobs[t.id] = j

	response.JobId = t.id
	return nil
}

func (s *Service) GetRecognitionStatus(ctx context.Context, request *rms_speech.GetRecognitionStatusRequest, response *rms_speech.GetRecognitionStatusResponse) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	j, ok := s.jobs[request.JobId]
	if !ok {
		return errors.New("job not found")
	}

	status := j.receipt.Status()
	response.Status = convertStatus(status)
	if status == worker.Done {
		response.RecognizedText = j.t.recognizedText
		delete(s.jobs, request.JobId)
	}

	return nil
}

func (s *Service) StopRecognition(ctx context.Context, request *rms_speech.StopRecognitionRequest, empty *emptypb.Empty) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	j, ok := s.jobs[request.JobId]
	if !ok {
		return errors.New("job not found")
	}

	j.receipt.Cancel()
	delete(s.jobs, request.JobId)

	return nil
}
