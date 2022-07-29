package api

// TokeninfoIntrospectResponse is a type for the response for an introspect request
type TokeninfoIntrospectResponse struct {
	Valid     bool        `json:"valid"`
	TokenType string      `json:"token_type"`
	Token     UsedMytoken `json:"token"`
}

// TokeninfoHistoryResponse is a type for the response for an history request
type TokeninfoHistoryResponse struct {
	EventHistory EventHistory     `json:"events"`
	TokenUpdate  *MytokenResponse `json:"token_update,omitempty"`
}

// TokeninfoSubtokensResponse is a type for the response for an tree request
type TokeninfoSubtokensResponse struct {
	Tokens      MytokenEntryTree `json:"mytokens"`
	TokenUpdate *MytokenResponse `json:"token_update,omitempty"`
}

// TokeninfoListResponse is a type for the response for an list request
type TokeninfoListResponse struct {
	Tokens      []MytokenEntryTree `json:"mytokens"`
	TokenUpdate *MytokenResponse   `json:"token_update,omitempty"`
}
