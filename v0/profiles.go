package api

import (
	"encoding/json"
)

type Profile struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Payload json.RawMessage `json:"payload"`
}
