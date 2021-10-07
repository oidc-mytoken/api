package api

// GrantTypeRequest is a request object for grant type requests at the grant type settings endpoint. The same GrantTypeRequest struct is used for Enable and Disable requests
type GrantTypeRequest struct {
	GrantType string `json:"grant_type" form:"grant_type" xml:"grant_type"`
	Mytoken   string `json:"mytoken" form:"mytoken" xml:"mytoken"`
}
