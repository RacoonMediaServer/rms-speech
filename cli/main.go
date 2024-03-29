package main

import (
	"context"
	"fmt"
	rms_speech "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-speech"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"os"
	"time"
)

func main() {
	var input string
	service := micro.NewService(
		micro.Name("rms-speech.client"),
		micro.Flags(
			&cli.StringFlag{
				Name:        "input",
				Usage:       "Audio file",
				Required:    true,
				Destination: &input,
			},
		),
	)
	service.Init()

	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	client := rms_speech.NewSpeechService("rms-speech", service.Client())
	req := rms_speech.StartRecognitionRequest{
		Data:        data,
		ContentType: "",
		TimeoutSec:  60,
	}

	resp, err := client.StartRecognition(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	for {
		status, err := client.GetRecognitionStatus(context.Background(), &rms_speech.GetRecognitionStatusRequest{JobId: resp.JobId})
		if err != nil {
			panic(err)
		}
		fmt.Println("Status: ", status.Status)
		if status.Status == rms_speech.GetRecognitionStatusResponse_Done {
			fmt.Println("Result:", status.RecognizedText)
			break
		}
		if status.Status == rms_speech.GetRecognitionStatusResponse_Failed {
			break
		}
		<-time.After(1 * time.Second)
	}
}
