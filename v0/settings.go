package api

// SettingsMetaData is the metadata related to settings
type SettingsMetaData struct {
	GrantTypeEndpoint string `json:"grant_type_endpoint,omitempty"`
	EmailEndpoint     string `json:"email_endpoint,omitempty"`
	TagsEndpoint      string `json:"tags_endpoint,omitempty"`
}
