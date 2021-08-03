package api

// CreateTransferCodeRequest is a request to create a new transfer code from an existing mytoken
type CreateTransferCodeRequest struct {
	Mytoken string `json:"mytoken"`
}

// ExchangeTransferCodeRequest is a request to exchange a transfer code for the mytoken
type ExchangeTransferCodeRequest struct {
	GrantType    string `json:"grant_type"`
	TransferCode string `json:"transfer_code"`
}
