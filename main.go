package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/revett/sepias/internal/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out: os.Stdout,
		},
	)

	root := cmd.Root()
	root.AddCommand(cmd.Completion())

	if err := root.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
