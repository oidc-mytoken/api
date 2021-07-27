package api

// MytokenFromMytokenRequest is a request to create a new Mytoken from an existing Mytoken
type MytokenFromMytokenRequest struct {
	GeneralMytokenRequest
	Mytoken                      string `json:"mytoken"`
	FailOnRestrictionsNotTighter bool   `json:"error_on_restrictions"`
}

// GeneralMytokenRequest is a type that holds all the information that all mytoken requests have in common. It should
// not be used directly as a request object
type GeneralMytokenRequest struct {
	Issuer               string       `json:"oidc_issuer"`
	GrantType            string       `json:"grant_type"`
	Restrictions         Restrictions `json:"restrictions"`
	Capabilities         Capabilities `json:"capabilities"`
	SubtokenCapabilities Capabilities `json:"subtoken_capabilities"`
	Name                 string       `json:"name"`
	ResponseType         string       `json:"response_type"`
	MaxTokenLen          int          `json:"max_token_len"`
	Rotation             *Rotation    `json:"rotation"`
}
