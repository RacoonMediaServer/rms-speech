package main

import (
	"fmt"
	rms_speech "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-speech"
	"github.com/RacoonMediaServer/rms-packages/pkg/worker"
	"github.com/RacoonMediaServer/rms-speech/internal/config"
	"github.com/RacoonMediaServer/rms-speech/internal/service/speech"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"time"

	// Plugins
	_ "github.com/go-micro/plugins/v4/registry/etcd"
)

var Version = "v0.0.0"

const serviceName = "rms-speech"

func main() {
	logger.Infof("%s %s", serviceName, Version)
	defer logger.Info("DONE.")

	useDebug := false

	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(Version),
		micro.Flags(
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"debug"},
				Usage:       "debug log level",
				Value:       false,
				Destination: &useDebug,
			},
		),
	)

	service.Init(
		micro.Action(func(context *cli.Context) error {
			configFile := fmt.Sprintf("/etc/rms/%s.json", serviceName)
			if context.IsSet("config") {
				configFile = context.String("config")
			}
			return config.Load(configFile)
		}),
	)

	cfg := config.Config()
	if useDebug || cfg.Debug.Verbose {
		_ = logger.Init(logger.WithLevel(logger.DebugLevel))
	}

	workers := worker.New(cfg.Workers, time.Duration(cfg.MaxJodDuration)*time.Minute)
	speechService := &speech.Service{
		Workers: workers,
	}

	// регистрируем хендлеры
	if err := rms_speech.RegisterSpeechHandler(service.Server(), speechService); err != nil {
		logger.Fatalf("Register service failed: %s", err)
	}

	if err := service.Run(); err != nil {
		logger.Fatalf("Run service failed: %s", err)
	}
	workers.Stop()
}
