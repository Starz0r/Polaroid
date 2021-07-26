package objstore

import (
	"os"
	"strconv"

	"github.com/minio/minio-go/v6"
)

var SESSION *minio.Client

func Login() error {
	err := *new(error)

	useSsl := true
	if os.Getenv("S3_USE_SSL") == "" {
		useSsl = false
	} else {
		useSsl, err = strconv.ParseBool(os.Getenv("S3_USE_SSL"))
		if err != nil {
			return err
		}
	}

	SESSION, err = minio.New(os.Getenv("S3_ENDPOINT"),
		os.Getenv("S3_ACCESS_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		useSsl,
	)
	return err
}
