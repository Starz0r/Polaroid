package objstore

import (
	"io"
	"os"
	"strconv"

	"github.com/minio/minio-go/v6"
	"github.com/spidernest-go/logger"
)

var (
	BUCKET          = os.Getenv("S3_BUCKET")
	UPLOAD_SIZE_MAX = os.Getenv("S3_UPLOAD_SIZE_MAX")
)

func Upload(f io.Reader, fname string, acl string) error {
	i, err := strconv.ParseInt(UPLOAD_SIZE_MAX, 10, 64)
	if err != nil {
		logger.Panic().Err(err).Msg("Upload size was not a valid int64.")
	}

	_, err = SESSION.PutObject(BUCKET,
		fname,
		f,
		i,
		minio.PutObjectOptions{
			ContentEncoding: "brotli",
		})
	return err
}
