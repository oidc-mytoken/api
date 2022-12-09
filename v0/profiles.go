package api

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Profile struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Payload json.RawMessage `json:"payload"`
}

type IncludedProfiles []string

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *IncludedProfiles) UnmarshalJSON(bytes []byte) error {
	var s string
	var arr []string
	if err := json.Unmarshal(bytes, &s); err != nil {
		err = errors.WithStack(json.Unmarshal(bytes, &arr))
		*i = arr
		return err
	} else {
		*i = NewIncludedProfiles(s)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (i IncludedProfiles) MarshalJSON() ([]byte, error) {
	if len(i) == 1 {
		return json.Marshal(i[0])
	}
	return json.Marshal([]string(i))
}

// NewIncludedProfiles creates a new IncludedProfiles var from the past profile names
func NewIncludedProfiles(profiles ...string) IncludedProfiles {
	return profiles
}
