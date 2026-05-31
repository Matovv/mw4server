package network

import "encoding/json"

type Packet struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
