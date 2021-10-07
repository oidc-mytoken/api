package api

type SSHKeyAddRequest struct {
	Mytoken              string       `json:"mytoken" form:"mytoken" xml:"mytoken"`
	SSHKey               string       `json:"ssh_key" form:"ssh_key" xml:"ssh_key"`
	Name                 string       `json:"name" form:"name" xml:"name"`
	Restrictions         Restrictions `json:"restrictions" form:"restrictions" xml:"restrictions"`
	Capabilities         Capabilities `json:"capabilities" form:"capabilities" xml:"capabilities"`
	SubtokenCapabilities Capabilities `json:"subtoken_capabilities" form:"subtoken_capabilities" xml:"subtoken_capabilities"`
}
