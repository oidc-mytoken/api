package api

type SSHInfoResponse struct {
	GrantEnabled bool             `json:"grant_enabled"`
	SSHKeyInfo   []SSHKeyInfo     `json:"ssh_keys"`
	TokenUpdate  *MytokenResponse `json:"token_update,omitempty"`
}

type SSHKeyInfo struct {
	Name       string `json:"name,omitempty"`
	SSHKey     string `json:"ssh_key,omitempty"`      // One of SSHKey and SSHKeyHash MUST be given
	SSHKeyHash string `json:"ssh_key_hash,omitempty"` // One of SSHKey and SSHKeyHash MUST be given
	Created    int64  `json:"created"`
	LastUsed   *int64 `json:"last_used,omitempty"`
}
