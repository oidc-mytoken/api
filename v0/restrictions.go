package api

// Restrictions is a slice of Restriction
type Restrictions []*Restriction

// Restriction describes a token usage restriction
type Restriction struct {
	NotBefore        int64            `json:"nbf,omitempty"`
	ExpiresAt        int64            `json:"exp,omitempty"`
	Scope            string           `json:"scope,omitempty"`
	Audiences        []string         `json:"audience,omitempty"`
	Hosts            []string         `json:"hosts,omitempty"`
	GeoIPAllow       []string         `json:"geoip_allow,omitempty"`
	GeoIPDisallow    []string         `json:"geoip_disallow,omitempty"`
	UsagesAT         *int64           `json:"usages_AT,omitempty"`
	UsagesOther      *int64           `json:"usages_other,omitempty"`
	IncludedProfiles IncludedProfiles `json:"include,omitempty"`
}

// UsedRestriction is a type for a restriction that has been used and additionally has information how often it has been used
type UsedRestriction struct {
	Restriction     `json:",inline"`
	UsagesATDone    *int64 `json:"usages_AT_done,omitempty"`
	UsagesOtherDone *int64 `json:"usages_other_done,omitempty"`
}
