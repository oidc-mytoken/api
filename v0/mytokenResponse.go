package api

// MytokenResponse is a response to a mytoken request
type MytokenResponse struct {
	Mytoken      string           `json:"mytoken,omitempty"`
	MytokenType  string           `json:"mytoken_type"`
	TransferCode string           `json:"transfer_code,omitempty"`
	MOMID        string           `json:"mom_id,omitempty"`
	ExpiresIn    uint64           `json:"expires_in,omitempty"`
	Restrictions Restrictions     `json:"restrictions,omitempty"`
	Capabilities Capabilities     `json:"capabilities,omitempty"`
	Rotation     *Rotation        `json:"rotation,omitempty"`
	TokenUpdate  *MytokenResponse `json:"token_update,omitempty"`
}
