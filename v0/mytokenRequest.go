package api

// MytokenFromMytokenRequest is a request to create a new Mytoken from an existing Mytoken
type MytokenFromMytokenRequest struct {
	GeneralMytokenRequest
	Mytoken                      string `json:"mytoken"`
	FailOnRestrictionsNotTighter bool   `json:"error_on_restrictions,omitempty"`
}

// GeneralMytokenRequest is a type that holds all the information that all mytoken requests have in common. It should
// not be used directly as a request object
type GeneralMytokenRequest struct {
	Issuer           string           `json:"oidc_issuer,omitempty"`
	GrantType        string           `json:"grant_type,omitempty"`
	Restrictions     Restrictions     `json:"restrictions,omitempty"`
	Capabilities     Capabilities     `json:"capabilities,omitempty"`
	Name             string           `json:"name,omitempty"`
	ResponseType     string           `json:"response_type,omitempty"`
	MaxTokenLen      int              `json:"max_token_len,omitempty"`
	Rotation         *Rotation        `json:"rotation,omitempty"`
	ApplicationName  string           `json:"application_name,omitempty"`
	IncludedProfiles IncludedProfiles `json:"include,omitempty"`
}
