package api

type SSHKeyDeleteRequest struct {
	Mytoken           string `json:"mytoken" form:"mytoken" xml:"mytoken"`
	SSHKey            string `json:"ssh_key" form:"ssh_key" xml:"ssh_key"`
	SSHKeyFingerprint string `json:"ssh_key_fp" form:"ssh_key_fp" xml:"ssh_key_fp"`
}
