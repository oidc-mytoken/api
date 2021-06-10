package api

// TokenInfoRequest is a type for requests to the tokeninfo endpoint
type TokenInfoRequest struct {
	Action  string `json:"action"`
	Mytoken string `json:"mytoken"`
}
