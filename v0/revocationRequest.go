package api

// RevocationRequest holds the information for a token revocation request
type RevocationRequest struct {
	Token      string `json:"token"`
	MOMID      string `json:"mom_id,omitempty"`
	Recursive  bool   `json:"recursive,omitempty"`
	OIDCIssuer string `json:"oidc_issuer,omitempty"`
}
