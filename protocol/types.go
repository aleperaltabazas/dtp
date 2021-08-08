package protocol

import (
	"encoding/json"
)

type Message struct {
	Code   string
	Source string
	Body   []byte
}

func (m * Message) Deserialize(as interface{}) error {
	return json.Unmarshal(m.Body[:], as)
}
