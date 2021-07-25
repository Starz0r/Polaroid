package main

import (
	"github.com/spidernest-go/logger"

	"github.com/Starz0r/Polaroid/src/auth"
	"github.com/Starz0r/Polaroid/src/crypto"
	"github.com/Starz0r/Polaroid/src/database"
)

func main() {
	logger.Info().Msg("Starting up.")

	err := database.Connect()
	if err != nil {
		logger.Fatal().Err(err).Msg("Database module failed to initialize.")
	}
	logger.Info().Msg("Connected to the database successfully.")

	err = crypto.InitRandomPool()
	if err != nil {
		logger.Fatal().Err(err).Msg("Cryptography module failed to initialize.")
	}
	logger.Info().Msg("Randomness pool has been successfully created.")

	err = auth.GetConfiguration()
	if err != nil {
		logger.Fatal().Err(err).Msg("Authentication module failed to initialize.")
	}
	logger.Info().Msg("OpenID Connect server has been configured for use.")
}
