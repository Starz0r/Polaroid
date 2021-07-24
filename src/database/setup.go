package database

import (
	"bytes"
	"database/sql"
	"io"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	// get env vars
	host := os.Getenv("PQ_HOST")
	database := os.Getenv("PQ_DB")
	port := os.Getenv("PQ_PORT")
	user := os.Getenv("PQ_USER")
	pass := os.Getenv("PQ_PASS")
	sslmode := os.Getenv("PQ_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	// constuct the url
	buf := bytes.NewBufferString("postgres://")

	buf.WriteString(user)
	buf.WriteString(":")

	buf.WriteString(pass)
	buf.WriteString("@")

	buf.WriteString(host)
	buf.WriteString(":")

	buf.WriteString(port)
	buf.WriteString("/")

	buf.WriteString(database)
	buf.WriteString("?sslmode=")

	buf.WriteString(sslmode)

	url, err := buf.ReadString(0)
	if err != nil && err != io.EOF {
		return err
	}

	// connect to database
	DB, err = sql.Open("postgres", url)
	if err != nil {
		return err
	}

	return nil
}
