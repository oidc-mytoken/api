package api

// AllTokeninfoActions holds all defined TokenInfo strings
var AllTokeninfoActions = [...]string{TokeninfoActionIntrospect, TokeninfoActionEventHistory,
	TokeninfoActionSubtokens, TokeninfoActionListMytokens}

// TokeninfoActions
const (
	TokeninfoActionIntrospect   = "introspect"
	TokeninfoActionEventHistory = "event_history"
	TokeninfoActionSubtokens    = "subtokens"
	TokeninfoActionListMytokens = "list_mytokens"
)
