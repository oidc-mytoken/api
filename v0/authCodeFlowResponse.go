package api

// AuthCodeFlowResponse is the response to an authorization code flow request
type AuthCodeFlowResponse struct {
	ConsentURI string `json:"consent_uri"`
	PollingInfo
}

// PollingInfo holds all response information about polling codes
type PollingInfo struct {
	PollingCode          string `json:"polling_code,omitempty"`
	PollingCodeExpiresIn int64  `json:"expires_in,omitempty"`
	PollingInterval      int64  `json:"interval,omitempty"`
}
