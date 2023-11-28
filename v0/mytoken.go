package api

import (
	"encoding/json"
	"fmt"
)

// Mytoken is a mytoken Mytoken
type Mytoken struct {
	Version      TokenVersion `json:"ver"`
	TokenType    string       `json:"token_type"`
	Issuer       string       `json:"iss"`
	Subject      string       `json:"sub"`
	ExpiresAt    int64        `json:"exp,omitempty"`
	NotBefore    int64        `json:"nbf"`
	IssuedAt     int64        `json:"iat"`
	AuthTime     int64        `json:"auth_time"`
	ID           string       `json:"jti"`
	SeqNo        uint64       `json:"seq_no"`
	Name         string       `json:"name,omitempty"`
	Audience     string       `json:"aud"`
	OIDCSubject  string       `json:"oidc_sub"`
	OIDCIssuer   string       `json:"oidc_iss"`
	Restrictions Restrictions `json:"restrictions,omitempty"`
	Capabilities Capabilities `json:"capabilities"`
	Rotation     *Rotation    `json:"rotation,omitempty"`
}

// TokenVer is the current Mytoken TokenVersion
var TokenVer = TokenVersion{
	Major: 0,
	Minor: 7,
}

// TokenType is a constant that can be used for identifying mytokens and to distinguish them from OIDC tokens
const TokenType = "mytoken"

// MinShortTokenLen is the minimum length of a short token
const MinShortTokenLen = 32

// UsedMytoken is a type for a Mytoken that has been used, it additionally has information how often it has been used
type UsedMytoken struct {
	Mytoken      `json:",inline"`
	Restrictions []UsedRestriction `json:"restrictions,omitempty"`
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

func (v TokenVersion) AtLeast(required TokenVersion) bool {
	if v.Major > required.Major {
		return true
	}
	if v.Major < required.Major {
		return false
	}
	// major version equal
	return v.Minor >= required.Minor
}
func (v TokenVersion) Before(before TokenVersion) bool {
	return !v.AtLeast(before)
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
