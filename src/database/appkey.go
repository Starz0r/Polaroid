package database

import (
	"errors"
	"time"

	"github.com/Starz0r/Polaroid/src/crypto"
)

type AppKey struct {
	ID          uint64    `db:"id" json:"id"`
	User        string    `db:"user" json:"user"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	Key         string    `db:"key" json:"key"`
}

func (ak *AppKey) New(u string) error {
	appkeys := DB.Collection("app_keys")
	ak = new(AppKey)

	if u == "" {
		return errors.New("username field cannot be empty")
	}

	ak.User = u
	ak.Key = crypto.String(128)

	_, err := appkeys.Insert(ak)

	return err
}

func (ak *AppKey) KeyExists() (bool, error) {
	appkeys := DB.Collection("app_keys")
	rs := appkeys.Find().Where("key = ", ak.Key)

	err := rs.One(ak)
	if err != nil {
		return false, err
	}
	return true, nil
}
