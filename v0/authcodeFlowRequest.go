package api

// AuthCodeFlowRequest holds a authorization code flow request
type AuthCodeFlowRequest struct {
	OIDCFlowRequest
	ClientType  string `json:"client_type"`
	RedirectURI string `json:"redirect_uri"`
}

// OIDCFlowRequest holds the request for an OIDC Flow request
type OIDCFlowRequest struct {
	GeneralMytokenRequest
	OIDCFlow string `json:"oidc_flow"`
}

// Client types
const (
	ClientTypeWeb    = "web"
	ClientTypeNative = "native"
)
