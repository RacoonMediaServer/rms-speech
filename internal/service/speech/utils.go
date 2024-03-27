package speech

import (
	rms_speech "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-speech"
	"github.com/RacoonMediaServer/rms-packages/pkg/worker"
)

func convertStatus(status worker.Status) rms_speech.GetRecognitionStatusResponse_Status {
	switch status {
	case worker.Pending:
		return rms_speech.GetRecognitionStatusResponse_Pending
	case worker.Active:
		return rms_speech.GetRecognitionStatusResponse_Processing
	case worker.Failed:
		return rms_speech.GetRecognitionStatusResponse_Failed
	case worker.Done:
		return rms_speech.GetRecognitionStatusResponse_Done
	default:
		return rms_speech.GetRecognitionStatusResponse_Failed
	}
}
