package utils

import (
	"github.com/codehack/scrypto"
	"github.com/google/uuid"
)

// HashPassword hash password using scrypto hash method
func HashPassword(password string) string {
	h, _ := scrypto.Hash(password)
	return h
}

// ComparePassword compare password and hash
func ComparePassword(p, h string) bool {
	return scrypto.Compare(p, h)
}

// MustUUID must generate a uuid
func MustUUID() string {
	uu, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return uu.String()
}
