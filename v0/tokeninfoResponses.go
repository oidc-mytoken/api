package api

// TokeninfoIntrospectResponse is a type for the response for an introspect request
type TokeninfoIntrospectResponse struct {
	Valid     bool        `json:"valid"`
	TokenType string      `json:"token_type"`
	Token     UsedMytoken `json:"token"`
	MOMID     string      `json:"mom_id"`
}

// TokeninfoHistoryResponse is a type for the response for a history request
type TokeninfoHistoryResponse struct {
	EventHistory
	TokenUpdate *MytokenResponse `json:"token_update,omitempty"`
}

// TokeninfoSubtokensResponse is a type for the response for a tree request
type TokeninfoSubtokensResponse struct {
	Tokens      MytokenEntryTree `json:"mytokens"`
	TokenUpdate *MytokenResponse `json:"token_update,omitempty"`
}

// TokeninfoListResponse is a type for the response for a list request
type TokeninfoListResponse struct {
	Tokens      []MytokenEntryTree `json:"mytokens"`
	TokenUpdate *MytokenResponse   `json:"token_update,omitempty"`
}
