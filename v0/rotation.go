package api

import (
	"database/sql/driver"
	"encoding/json"
)

// Rotation is a type describing how a mytoken might be rotated
type Rotation struct {
	OnAT             bool             `json:"on_AT,omitempty"`
	OnOther          bool             `json:"on_other,omitempty"`
	Lifetime         uint64           `json:"lifetime,omitempty"`
	AutoRevoke       bool             `json:"auto_revoke,omitempty"`
	IncludedProfiles IncludedProfiles `json:"include,omitempty"`
}

// Scan implements the sql.Scanner interface.
func (r *Rotation) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	val := src.([]uint8)
	err := json.Unmarshal(val, &r)
	return err
}

// Value implements the driver.Valuer interface
func (r Rotation) Value() (driver.Value, error) {
	if !r.OnAT && !r.OnOther {
		return nil, nil
	}
	return json.Marshal(r)
}
