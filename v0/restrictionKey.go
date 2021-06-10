package api

// AllRestrictionKeys holds all the RestrictionKeys
var AllRestrictionKeys = [...]string{RestrictionKeyNotBefore, RestrictionKeyExpiresAt, RestrictionKeyScope,
	RestrictionKeyAudiences, RestrictionKeyIPs, RestrictionKeyGeoIPAllow, RestrictionKeyGeoIPDisallow,
	RestrictionKeyUsagesAT, RestrictionKeyUsagesOther}

// RestrictionKeys
const (
	RestrictionKeyNotBefore     = "nbf"
	RestrictionKeyExpiresAt     = "exp"
	RestrictionKeyScope         = "scope"
	RestrictionKeyAudiences     = "audience"
	RestrictionKeyIPs           = "ip"
	RestrictionKeyGeoIPAllow    = "geoip_allow"
	RestrictionKeyGeoIPDisallow = "geoip_disallow"
	RestrictionKeyUsagesAT      = "usages_AT"
	RestrictionKeyUsagesOther   = "usages_other"
)
