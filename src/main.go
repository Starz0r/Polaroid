package main

import (
	"github.com/spidernest-go/logger"

	"github.com/Starz0r/Polaroid/src/crypto"
	"github.com/Starz0r/Polaroid/src/database"
)

func main() {
	logger.Info().Msg("Starting up.")

	err := database.Connect()
	if err != nil {
		logger.Fatal().Err(err).Msg("Database module failed to initialize.")
	}
	logger.Info().Msg("Successfully connected to the database.")

	err = crypto.InitRandomPool()
	if err != nil {
		logger.Fatal().Err(err).Msg("Cryptography module failed to initialize.")
	}
	logger.Info().Msg("Randomness pool has been successfully created.")
}
