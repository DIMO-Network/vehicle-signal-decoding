package main

import (
	"context"
	"log"
	"os"

	"github.com/DIMO-Network/shared"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/api"
	"github.com/DIMO-Network/vehicle-signal-decoding/internal/config"
	"github.com/rs/zerolog"
)

func main() {

	gitSha1 := os.Getenv("GIT_SHA1")
	ctx := context.Background()
	arg := ""

	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	settings, err := shared.LoadConfig[config.Settings]("settings.yaml")
	if err != nil {
		log.Fatal("could not load settings: $s", err)
	}
	level, err := zerolog.ParseLevel(settings.LogLevel)
	if err != nil {
		log.Fatal("could not parse log level: $s", err)
	}
	logger := zerolog.New(os.Stdout).Level(level).With().
		Timestamp().
		Str("app", settings.ServiceName).
		Str("git-sha1", gitSha1).
		Logger()

	switch arg {
	case "migrate":
		migrateDatabase(ctx, logger, &settings, os.Args)
	default:
		api.Run(ctx, logger, &settings)
	}
}
