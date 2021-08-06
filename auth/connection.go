package auth

import (
	bytes2 "bytes"
	"crypto/sha256"
	"github.com/aleperaltabazas/dtp/config"
)

func Passphrase() []byte {
	passphrase := config.Config().GetString("connection.passphrase")
	// TODO: use a better key to hash the passphrase
	bytes := sha256.Sum256([]byte(passphrase))
	return bytes[:]
}

func Authenticate(bytes []byte) bool {
	return bytes2.Compare(bytes, Passphrase()) == 0
}
