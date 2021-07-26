package database

import (
	"os"

	"github.com/spidernest-go/db/lib/sqlbuilder"
	"github.com/spidernest-go/db/postgresql"
)

var DB sqlbuilder.Database

func Connect() error {
	// constuct the url
	opts := make(map[string]string)

	sslmode := os.Getenv("PQ_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	opts["sslmode"] = sslmode

	url := postgresql.ConnectionURL{
		Host:     os.Getenv("PQ_HOST") + ":" + os.Getenv("PQ_PORT"),
		Database: os.Getenv("PQ_DB"),
		User:     os.Getenv("PQ_USER"),
		Password: os.Getenv("PQ_PASS"),
		Options:  opts,
	}

	// connect to database
	err := *new(error)
	DB, err = postgresql.Open(url)
	if err != nil {
		return err
	}

	return nil
}
