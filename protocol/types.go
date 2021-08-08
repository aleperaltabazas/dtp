package protocol

import (
	"bytes"
	"encoding/gob"
)

type Message struct {
	Code   string
	Source string
	Body   []byte
}

func (m * Message) Deserialize(as interface{}) error {
	buf := bytes.NewBuffer(m.Body)
	dec := gob.NewDecoder(buf)

	return dec.Decode(&as)
}
