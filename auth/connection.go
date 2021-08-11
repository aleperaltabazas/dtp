package auth

import (
	"bytes"
	"crypto/sha256"
	"github.com/aleperaltabazas/dtp/config"
)

func Passphrase() []byte {
	passphrase := config.Config.Passphrase
	b := sha256.Sum256([]byte(passphrase))
	return b[:]
}

func Authenticate(b []byte) bool {
	return bytes.Compare(b, Passphrase()) == 0
}
