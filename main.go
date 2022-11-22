package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/digvijay17july/SqlDbToCSV/api"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Starting the DB Set up")
	api.Connect()
	log.Info().Msg("Starting the Http Server")
	api.Start()

}
