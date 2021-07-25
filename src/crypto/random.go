package crypto

import (
	srand "crypto/rand"
	"math/rand"
)

var randpool rand.Source

func InitRandomPool() error {
	seed, err := srand.Prime(srand.Reader, 256)
	if err != nil {
		return err
	}
	randpool = rand.NewSource(seed.Int64())
	return nil
}
