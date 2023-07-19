package api

const (
	MOMIDValueThis     = "this"
	MOMIDValueChildren = "children"
)

// TokenInfoRequest is a type for requests to the tokeninfo endpoint
type TokenInfoRequest struct {
	Action  string   `json:"action"`
	Mytoken string   `json:"mytoken"`
	MOMIDs  []string `json:"mom_ids,omitempty"`
}
