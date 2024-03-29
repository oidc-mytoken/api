package api

type SSHKeyAddRequest struct {
	Mytoken      string       `json:"mytoken" form:"mytoken" xml:"mytoken"`
	SSHKey       string       `json:"ssh_key" form:"ssh_key" xml:"ssh_key"`
	Name         string       `json:"name" form:"name" xml:"name"`
	Restrictions Restrictions `json:"restrictions" form:"restrictions" xml:"restrictions"`
	Capabilities Capabilities `json:"capabilities" form:"capabilities" xml:"capabilities"`
	GrantType    string       `json:"grant_type" form:"grant_type" xml:"grant_type"`
}
