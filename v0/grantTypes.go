package api

// AllGrantTypes holds all the GrantTypes
var AllGrantTypes = [...]string{GrantTypeMytoken, GrantTypeOIDCFlow, GrantTypePollingCode, GrantTypeTransferCode, GrantTypeSSH}

// GrantTypes
const (
	GrantTypeMytoken      = "mytoken"
	GrantTypeOIDCFlow     = "oidc_flow"
	GrantTypePollingCode  = "polling_code"
	GrantTypeTransferCode = "transfer_code"
	GrantTypeSSH          = "ssh"
)
