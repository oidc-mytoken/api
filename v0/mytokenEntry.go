package api

// MytokenEntry holds the information of a MytokenEntry as stored in the
// database
type MytokenEntry struct {
	// The "Manage-Other-Mytokens-ID (MOMID) is used in requests to manage other mytokens than the one used for
	// authorization
	MOMID          string `json:"mom_id"`
	Name           string `json:"name,omitempty"`
	CreatedAt      int64  `json:"created"`
	ClientMetaData `json:",inline"`
}

// MytokenEntryTree is a tree of MytokenEntry
type MytokenEntryTree struct {
	Token    MytokenEntry       `json:"token"`
	Children []MytokenEntryTree `json:"children,omitempty"`
}
