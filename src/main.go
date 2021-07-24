package main

import (
	"github.com/spidernest-go/logger"

	"github.com/Starz0r/Polaroid/src/database"
)

func main() {
	logger.Info().Msg("Starting up.")

	err := database.Connect()
	if err != nil {
		logger.Panic().Err(err).Msg("Database could not be accessed.")
	}
	logger.Info().Msg("Successfully connected to the database.")

}
