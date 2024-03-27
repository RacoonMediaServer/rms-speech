package speech

import (
	"context"
	rms_speech "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-speech"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	Workers Workers
}

func (s Service) StartRecognition(ctx context.Context, request *rms_speech.StartRecognitionRequest, response *rms_speech.StartRecognitionResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetRecognitionStatus(ctx context.Context, request *rms_speech.GetRecognitionStatusRequest, response *rms_speech.GetRecognitionStatusResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) StopRecognition(ctx context.Context, request *rms_speech.StopRecognitionRequest, empty *emptypb.Empty) error {
	//TODO implement me
	panic("implement me")
}
