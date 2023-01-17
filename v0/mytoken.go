package api

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
	Minor: 6,
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
