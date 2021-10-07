package api

type GrantTypeInfoResponse struct {
	GrantTypes  []GrantTypeInfo  `json:"grant_types"`
	TokenUpdate *MytokenResponse `json:"token_update,omitempty"`
}

// GrantTypeInfo is a struct holding information indicating if a grant type is enabled or not
type GrantTypeInfo struct {
	GrantType string `json:"grant_type"`
	Enabled   bool   `json:"enabled"`
}
