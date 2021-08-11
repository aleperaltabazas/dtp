package tcp

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"github.com/aleperaltabazas/dtp/config"
	"io"
	"log"
	"net"
)

var cip cipher.Block

func init() {
	key := config.Config.CipherKey
	c, err := aes.NewCipher([]byte(key))

	if err != nil {
		log.Fatal(err)
	}

	cip = c
}

func encrypt(bytes []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(cip)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	c := gcm.Seal(nil, nonce, bytes, nil)
	return append(nonce, c...), err
}

func decrypt(bytes []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(cip)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(bytes) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := bytes[:nonceSize], bytes[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func Send(c *net.TCPConn, message interface{}) error {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(message)

	if err != nil {
		return err
	}

	bs, err := encrypt(buf.Bytes())

	if err != nil {
		return err
	}

	enc := gob.NewEncoder(c)
	return enc.Encode(bs)
}

func Receive(c *net.TCPConn, message interface{}) error {
	dec := gob.NewDecoder(c)
	var buf []byte

	err := dec.Decode(&buf)

	if err != nil {
		return err
	}

	plainText, err := decrypt(buf)

	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(bytes.NewBuffer(plainText))
	err = decoder.Decode(message)

	if err != nil {
		return err
	}

	return nil
}
