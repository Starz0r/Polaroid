package main

import (
	"github.com/spidernest-go/logger"

	"github.com/Starz0r/Polaroid/src/auth"
	"github.com/Starz0r/Polaroid/src/crypto"
	"github.com/Starz0r/Polaroid/src/database"
	"github.com/Starz0r/Polaroid/src/objstore"
	"github.com/Starz0r/Polaroid/src/routers"
)

func main() {
	logger.Info().Msg("Starting up.")

	err := database.Connect()
	if err != nil {
		logger.Fatal().Err(err).Msg("Database module failed to initialize.")
	}
	logger.Info().Msg("Connected to the database successfully.")

	err = database.Synchronize()
	if err != nil {
		logger.Fatal().Err(err).Msg("Database module failed to apply migrations.")
	}
	logger.Info().Msg("Database migrations have been applied.")

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

	err = objstore.Login()
	if err != nil {
		logger.Fatal().Err(err).Msg("Object Storage module failed to initialize.")
	}
	logger.Info().Msg("Using the Object Storage bucket.")

	routers.ListenAndServe()
}
