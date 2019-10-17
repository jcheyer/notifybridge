package main

import (
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("starting notifybridge...")

	h := new(handler)

	web := engine(h)
	if err := web.Run(":9090"); err != nil {
		log.Fatal().Err(err).Msg("webserver failed")
	}
}
