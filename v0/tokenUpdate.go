package api

// OnlyTokenUpdateResponse is a type for responses that do not contain any content except for a (
// possibly) updated mytoken
type OnlyTokenUpdateResponse struct {
	TokenUpdate *MytokenResponse `json:"token_update,omitempty"`
}
