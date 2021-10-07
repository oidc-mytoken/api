package api

// AuthCodeFlowRequest holds a authorization code flow request
type AuthCodeFlowRequest struct {
	OIDCFlowRequest
	RedirectType string `json:"redirect_type"`
}

// OIDCFlowRequest holds the request for an OIDC Flow request
type OIDCFlowRequest struct {
	GeneralMytokenRequest
	OIDCFlow string `json:"oidc_flow"`
}

// Redirect types
const (
	RedirectTypeWeb    = "web"
	RedirectTypeNative = "native"
)
