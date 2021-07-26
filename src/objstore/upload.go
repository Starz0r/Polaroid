package objstore

import (
	"io"
	"os"

	"github.com/minio/minio-go/v6"
)

var BUCKET = os.Getenv("S3_BUCKET")

func Upload(f io.Reader, fname string) error {
	_, err := SESSION.PutObject(BUCKET,
		"/"+fname,
		f,
		-1,
		minio.PutObjectOptions{
			ContentEncoding: "brotli",
		})
	return err
}
