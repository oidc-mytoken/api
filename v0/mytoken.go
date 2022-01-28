package api

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Mytoken is a mytoken Mytoken
type Mytoken struct {
	Version              TokenVersion `json:"ver"`
	TokenType            string       `json:"token_type"`
	Issuer               string       `json:"iss"`
	Subject              string       `json:"sub"`
	ExpiresAt            int64        `json:"exp,omitempty"`
	NotBefore            int64        `json:"nbf"`
	IssuedAt             int64        `json:"iat"`
	AuthTime             int64        `json:"auth_time"`
	ID                   string       `json:"jti"`
	SeqNo                uint64       `json:"seq_no"`
	Name                 string       `json:"name,omitempty"`
	Audience             string       `json:"aud"`
	OIDCSubject          string       `json:"oidc_sub"`
	OIDCIssuer           string       `json:"oidc_iss"`
	Restrictions         Restrictions `json:"restrictions,omitempty"`
	Capabilities         Capabilities `json:"capabilities"`
	SubtokenCapabilities Capabilities `json:"subtoken_capabilities,omitempty"`
	Rotation             *Rotation    `json:"rotation,omitempty"`
}

// TokenVer is the current Mytoken TokenVersion
var TokenVer = TokenVersion{
	Major: 0,
	Minor: 4,
}

// TokenType is constant that can be used for identifying mytokens and to distinguish them from OIDC tokens
const TokenType = "mytoken"

// UsedMytoken is a type for a Mytoken that has been used, it additionally has information how often it has been used
type UsedMytoken struct {
	Mytoken      `json:",inline"`
	Restrictions []UsedRestriction `json:"restrictions,omitempty"`
}

// Rotation is a type describing how a mytoken might be rotated
type Rotation struct {
	OnAT       bool   `json:"on_AT,omitempty"`
	OnOther    bool   `json:"on_other,omitempty"`
	Lifetime   uint64 `json:"lifetime,omitempty"`
	AutoRevoke bool   `json:"auto_revoke,omitempty"`
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
