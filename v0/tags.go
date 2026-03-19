package api

// Tag is a type representing a tag name
type Tag string

// CreateMytokenTag is used when specifying tags during mytoken creation
type CreateMytokenTag struct {
	Tag             Tag  `json:"tag"`
	IncludeChildren bool `json:"include_children,omitempty"`
}

// TagInfo provides information about a tag
type TagInfo struct {
	Tag   Tag    `json:"tag" db:"tag"`
	Color string `json:"color" db:"color"`
}

type MTTagInfo struct {
	TagInfo
	IncludeChildren bool `json:"include_children"`
}

// TagListingResponse is the response type that lists tags of a user
type TagListingResponse struct {
	Tags        []TagInfo        `json:"tags"`
	TokenUpdate *MytokenResponse `json:"token_update,omitempty"`
}

type AddTagToMytokenRequest struct {
	Tag             Tag    `json:"tag"`
	Mytoken         string `json:"mytoken,omitempty" form:"mytoken" xml:"mytoken"`
	MOMID           string `json:"mom_id,omitempty" form:"mom_id" xml:"mom_id"`
	IncludeChildren bool   `json:"include_children" form:"include_children" xml:"include_children"`
}
type RemoveTagFromMytokenRequest struct {
	Tag             Tag    `json:"tag"`
	Mytoken         string `json:"mytoken,omitempty" form:"mytoken" xml:"mytoken"`
	MOMID           string `json:"mom_id,omitempty" form:"mom_id" xml:"mom_id"`
	IncludeChildren bool   `json:"include_children" form:"include_children" xml:"include_children"`
}
