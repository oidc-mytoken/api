package api

import (
	"fmt"
)

// Error is an error object that is returned on the api when an error occurs
type Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// CombinedMessage returns an error string that combines the error with the error description
func (e Error) CombinedMessage() string {
	if e.ErrorDescription == "" {
		return e.Error
	}
	return fmt.Sprintf("%s: %s", e.Error, e.ErrorDescription)
}

// Predefined errors
var (
	ErrorUnknownIssuer            = Error{ErrorStrInvalidRequest, "The provided issuer is not supported"}
	ErrorStateMismatch            = Error{ErrorStrInvalidRequest, "State mismatched"}
	ErrorUnsupportedOIDCFlow      = Error{ErrorStrInvalidGrant, "Unsupported oidc_flow"}
	ErrorUnsupportedGrantType     = Error{ErrorStrInvalidGrant, "Unsupported grant_type"}
	ErrorBadTransferCode          = Error{ErrorStrInvalidToken, "Bad polling or transfer code"}
	ErrorTransferCodeExpired      = Error{ErrorStrExpiredToken, "polling or transfer code is expired"}
	ErrorAuthorizationPending     = Error{ErrorStrAuthorizationPending, ""}
	ErrorConsentDeclined          = Error{ErrorStrAccessDenied, "user declined consent"}
	ErrorNoRefreshToken           = Error{ErrorStrOIDC, "Did not receive a refresh token"}
	ErrorInsufficientCapabilities = Error{ErrorStrInsufficientCapabilities, "The provided token does not have the required capability for this operation"}
	ErrorUsageRestricted          = Error{ErrorStrUsageRestricted, "The restrictions of this token does not allow this usage"}
	ErrorNYI                      = Error{ErrorStrNYI, ""}
)

// Predefined OAuth2/OIDC errors
const (
	ErrorStrInvalidRequest       = "invalid_request"
	ErrorStrInvalidClient        = "invalid_client"
	ErrorStrInvalidGrant         = "invalid_grant"
	ErrorStrUnauthorizedClient   = "unauthorized_client"
	ErrorStrUnsupportedGrantType = "unsupported_grant_type"
	ErrorStrInvalidScope         = "invalid_scope"
	ErrorStrInvalidToken         = "invalid_token"
	ErrorStrInsufficientScope    = "insufficient_scope"
	ErrorStrExpiredToken         = "expired_token"
	ErrorStrAccessDenied         = "access_denied"
	ErrorStrAuthorizationPending = "authorization_pending"
)

// Additional Mytoken errors
const (
	ErrorStrInternal                 = "internal_server_error"
	ErrorStrOIDC                     = "oidc_error"
	ErrorStrNYI                      = "not_yet_implemented"
	ErrorStrInsufficientCapabilities = "insufficient_capabilities"
	ErrorStrUsageRestricted          = "usage_restricted"
)
