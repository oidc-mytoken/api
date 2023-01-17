package api

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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

// TokenVersion is a type for the mytoken version
type TokenVersion struct {
	Major int
	Minor int
}

const vFmt = "%d.%d"

// String returns a version string
func (v TokenVersion) String() string {
	return fmt.Sprintf(vFmt, v.Major, v.Minor)
}

// Version returns a version string
func (v TokenVersion) Version() string {
	return v.String()
}

// MarshalJSON implements the json.Marshaler interface
func (v TokenVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (v *TokenVersion) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	_, err := fmt.Sscanf(str, vFmt, &v.Major, &v.Minor)
	return err
}
