package main

import (
	"os"

	"example.com/go-inmem-loki/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	if os.Getenv("LOG_LEVEL") == "INFO" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Info().Str("Component", "main").Msgf("Process ID: %v", os.Getpid())
	cmd.Execute()
}
