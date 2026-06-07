package network

import "encoding/json"

type Packet struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func NewPacket(
	packetType string,
	payload any,
) ([]byte, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	packet := Packet{
		Type: packetType,
		Data: data,
	}
	return json.Marshal(packet)
}
