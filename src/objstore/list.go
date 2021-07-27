package objstore

import (
	"github.com/minio/minio-go/v6"
)

func IsNameConflicting(name string) bool {
	_, err := SESSION.GetObject(BUCKET, name, minio.GetObjectOptions{})
	if err != nil {
		return false
	}
	return true
}
