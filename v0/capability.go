package api

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

// Defined Capabilities
var (
	CapabilityAT = Capability{
		Name:        "AT",
		Description: "Allows obtaining OpenID Connect Access Tokens.",
	}
	CapabilityCreateMT = Capability{
		Name:        "create_mytoken",
		Description: "Allows to create a new mytoken.",
	}
	CapabilitySettings = Capability{
		Name:        "settings",
		Description: "Allows access to user settings.",
	}
	CapabilitySettingsRead = Capability{
		Name:        CapabilityReadOnlyPrefix + CapabilitySettings.Name,
		Description: "Allows read access to user settings.",
	}
	CapabilityGrants = Capability{
		Name:        CapabilitySettings.Name + ":grants",
		Description: "Allows access to user grants.",
	}
	CapabilityGrantsRead = Capability{
		Name:        CapabilityReadOnlyPrefix + CapabilityGrants.Name,
		Description: "Allows read access to user grants.",
	}
	CapabilitySSHGrant = Capability{
		Name:        CapabilityGrants.Name + ":ssh",
		Description: "Allows access to the ssh grant.",
	}
	CapabilitySSHGrantRead = Capability{
		Name:        CapabilityReadOnlyPrefix + CapabilitySSHGrant.Name,
		Description: "Allows read access to the ssh grant.",
	}
	CapabilityTokeninfo = Capability{
		Name:        "tokeninfo",
		Description: "Allows to obtain all information about this token.",
	}
	CapabilityTokeninfoIntrospect = Capability{
		Name:        CapabilityTokeninfo.Name + ":introspect",
		Description: "Allows to obtain basic information about this token.",
	}
	CapabilityTokeninfoHistory = Capability{
		Name:        CapabilityTokeninfo.Name + ":history",
		Description: "Allows to obtain the event history for this token.",
	}
	CapabilityTokeninfoSubtokens = Capability{
		Name:        CapabilityTokeninfo.Name + ":subtokens",
		Description: "Allows to list a subtoken-tree for this token.",
	}
	CapabilityListMT = Capability{
		Name:        "list_mytokens",
		Description: "Allows to list all mytokens.",
	}
)

// AllCapabilities holds all defined Capabilities
var AllCapabilities = Capabilities{
	CapabilityAT,
	CapabilityCreateMT,
	CapabilitySettings,
	CapabilitySettingsRead,
	CapabilityGrants,
	CapabilityGrantsRead,
	CapabilitySSHGrant,
	CapabilitySSHGrantRead,
	CapabilityTokeninfo,
	CapabilityTokeninfoIntrospect,
	CapabilityTokeninfoHistory,
	CapabilityTokeninfoSubtokens,
	CapabilityListMT,
}

func descriptionFor(name string) string {
	for _, c := range AllCapabilities {
		if strings.EqualFold(c.Name, name) {
			return c.Description
		}
	}
	return ""
}

// NewCapabilities casts a []string into Capabilities
func NewCapabilities(caps []string) (c Capabilities) {
	for _, cc := range caps {
		c = append(c, NewCapability(cc))
	}
	return
}

// NewCapability casts a string into a Capability
func NewCapability(name string) Capability {
	return Capability{
		Name:        name,
		Description: descriptionFor(name),
	}
}

// Strings returns a slice of strings for these capabilities
func (c Capabilities) Strings() (s []string) {
	for _, cc := range c {
		s = append(s, cc.Name)
	}
	return
}

// Capabilities is a slice of Capability
type Capabilities []Capability

// Capability is a capability string
type Capability struct {
	Name        string
	Description string
}

// MarshalJSON implements the json.Marshaler interface
func (c Capability) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (c *Capability) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*c = NewCapability(name)
	return nil
}

// Scan implements the sql.Scanner interface.
func (c *Capabilities) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	val := src.([]uint8)
	err := json.Unmarshal(val, &c)
	return err
}

// Value implements the driver.Valuer interface
func (c Capabilities) Value() (driver.Value, error) {
	if len(c) == 0 {
		return nil, nil
	}
	return json.Marshal(c)
}

// TightenCapabilities tightens two set of Capabilities into one new
func TightenCapabilities(a, b Capabilities) (res Capabilities) {
	if b == nil {
		return a
	}
	for _, bb := range b {
		if a.Has(bb) {
			res = append(res, bb)
		}
	}
	return
}

const CapabilityReadOnlyPrefix = "read@"

func (c Capability) parse() (parts []string, readOnly bool) {
	s := c.Name
	if strings.HasPrefix(s, CapabilityReadOnlyPrefix) {
		readOnly = true
		s = s[len(CapabilityReadOnlyPrefix):]
	}
	parts = strings.Split(s, ":")
	return
}

// Has checks if Capabilities slice contains the passed Capability
func (c Capabilities) Has(a Capability) bool {
	aParts, aReadOnly := a.parse()
	for _, cc := range c {
		cParts, cReadOnly := cc.parse()
		if cReadOnly && !aReadOnly {
			continue
		}
		if len(cParts) > len(aParts) {
			continue
		}
		notValid := false
		for i, cPart := range cParts {
			if cPart != aParts[i] {
				notValid = true
				break
			}
		}
		if !notValid {
			return true
		}
	}
	return false
}
