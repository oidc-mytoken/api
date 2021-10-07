package api

type SSHKeyDeleteRequest struct {
	Mytoken    string `json:"mytoken" form:"mytoken" xml:"mytoken"`
	SSHKey     string `json:"ssh_key" form:"ssh_key" xml:"ssh_key"`
	SSHKeyHash string `json:"ssh_key_hash" form:"ssh_key_hash" xml:"ssh_key_hash"`
}
