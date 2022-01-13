package api

// SSHKeyAddResponse is a type for the (first) response to an SSHKeyAddRequest
type SSHKeyAddResponse struct {
	AuthCodeFlowResponse
	TokenUpdate *MytokenResponse `json:"token_update,omitempty"`
}

// SSHKeyAddFinalResponse is a type for the final response for an SSHKeyAddRequest after the polling was successful
type SSHKeyAddFinalResponse struct {
	SSHUser       string `json:"ssh_user"`
	SSHHostConfig string `json:"ssh_host_config,omitempty"`
}
