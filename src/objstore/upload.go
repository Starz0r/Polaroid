package objstore

import (
	"io"
	"os"

	"github.com/minio/minio-go/v6"
)

var BUCKET = os.Getenv("S3_BUCKET")

func Upload(f io.Reader, fname string, acl string) error {
	metadata := map[string]string{"x-amz-acl": acl}
	_, err := SESSION.PutObject(BUCKET,
		fname,
		f,
		-1,
		minio.PutObjectOptions{
			ContentEncoding: "brotli",
			UserMetadata:    metadata,
		})
	return err
}
