package api

// AllRestrictionClaims holds all the RestrictionClaims
var AllRestrictionClaims = [...]string{
	RestrictionClaimNotBefore,
	RestrictionClaimExpiresAt,
	RestrictionClaimScope,
	RestrictionClaimAudiences,
	RestrictionClaimHosts,
	RestrictionClaimGeoIPAllow,
	RestrictionClaimGeoIPDisallow,
	RestrictionClaimUsagesAT,
	RestrictionClaimUsagesOther,
}

// RestrictionClaims
const (
	RestrictionClaimNotBefore     = "nbf"
	RestrictionClaimExpiresAt     = "exp"
	RestrictionClaimScope         = "scope"
	RestrictionClaimAudiences     = "audience"
	RestrictionClaimHosts         = "hosts"
	RestrictionClaimGeoIPAllow    = "geoip_allow"
	RestrictionClaimGeoIPDisallow = "geoip_disallow"
	RestrictionClaimUsagesAT      = "usages_AT"
	RestrictionClaimUsagesOther   = "usages_other"
)
