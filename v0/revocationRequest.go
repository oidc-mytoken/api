package api

// RevocationRequest holds the information for a token revocation request
type RevocationRequest struct {
	Token      string `json:"token"`
	Recursive  bool   `json:"recursive,omitempty"`
	OIDCIssuer string `json:"oidc_issuer,omitempty"`
}
