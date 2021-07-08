package api

// TokeninfoIntrospectResponse is a type for the response for an introspect request
type TokeninfoIntrospectResponse struct {
	Valid bool        `json:"valid"`
	Token UsedMytoken `json:"token"`
}

// TokeninfoHistoryResponse is a type for the response for an history request
type TokeninfoHistoryResponse struct {
	EventHistory EventHistory     `json:"events"`
	TokenUpdate  *MytokenResponse `json:"token_update,omitempty"`
}

// TokeninfoTreeResponse is a type for the response for an tree request
type TokeninfoTreeResponse struct {
	Tokens      MytokenEntryTree `json:"mytokens"`
	TokenUpdate *MytokenResponse `json:"token_update,omitempty"`
}

// TokeninfoListResponse is a type for the response for an list request
type TokeninfoListResponse struct {
	Tokens      []MytokenEntryTree `json:"mytokens"`
	TokenUpdate *MytokenResponse   `json:"token_update,omitempty"`
}
