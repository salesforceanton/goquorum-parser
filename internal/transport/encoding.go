package transport

import "encoding/json"

type ProtocolMessage any

func Marshal(msg ProtocolMessage) ([]byte, error) {
	return json.Marshal(msg)
}

func Unmarshal(data []byte, msg ProtocolMessage) error {
	return json.Unmarshal(data, msg)
}
