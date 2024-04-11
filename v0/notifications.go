package api

import (
	"encoding/json"
)

// Variables for different notification types
const (
	NotificationTypeMail      = "mail"
	NotificationTypeWebsocket = "ws"
	NotificationTypeICSInvite = "ics_invite"
)

// NotificationsCreateResponse is a type holding the response returned after a new notification was created
type NotificationsCreateResponse struct {
	ManagementCode string           `json:"management_code"`
	TokenUpdate    *MytokenResponse `json:"token_update,omitempty"`
}

// NotificationsCombinedResponse is a type holding a response including information about notifications and calendars
type NotificationsCombinedResponse struct {
	Notifications []NotificationInfo `json:"notifications,omitempty"`
	Calendars     []CalendarInfo     `json:"calendars,omitempty"`
	TokenUpdate   *MytokenResponse   `json:"token_update,omitempty"`
}

// SubscribeNotificationRequest is a type holding a request to subscribe / create a new notification
type SubscribeNotificationRequest struct {
	Mytoken             string               `json:"mytoken,omitempty" form:"mytoken" xml:"mytoken"`
	MOMID               string               `json:"mom_id,omitempty" form:"mom_id" xml:"mom_id"`
	NotificationType    string               `json:"notification_type" form:"notification_type" xml:"notification_type"`
	NotificationClasses []*NotificationClass `json:"notification_classes" form:"notification_classes" xml:"notification_classes"`
	IncludeChildren     bool                 `json:"include_children"`
	UserWide            bool                 `json:"user_wide"`
	Comment             string               `json:"comment,omitempty" form:"comment" xml:"comment"` // only for notification_type=ics_invite
}

// NotificationInfo is a type for holding information about a notification including the subscribed
// NotificationClasses and tokens
type NotificationInfo struct {
	NotificationInfoBase
	Classes          []*NotificationClass `json:"notification_classes"`
	SubscribedTokens []string             `json:"subscribed_tokens,omitempty"`
}

// NotificationInfoBase is a type for holding the base information about a notification
type NotificationInfoBase struct {
	NotificationID uint64 `db:"id" json:"notification_id"`
	Type           string `db:"type" json:"notification_type"`
	ManagementCode string `db:"management_code" json:"management_code"`
	WebSocketPath  string `db:"ws" json:"ws,omitempty"`
	UserWide       bool   `db:"user_wide" json:"user_wide"`
}

// ManagementCodeNotificationInfoResponse is the response to a request to obtain infos about a notification via a
// management code
// This response will contain the oidc issuer url,
// this allows the client to map the notification to an issuer and also to obtain a mytoken for the correct issuer in
// order to obtain additional information
type ManagementCodeNotificationInfoResponse struct {
	NotificationInfo
	OIDCIssuer string `json:"oidc_iss"`
}

// NotificationAddTokenRequest is a type holding a request to add a token a notification
type NotificationAddTokenRequest struct {
	Mytoken         string `json:"mytoken,omitempty" form:"mytoken" xml:"mytoken"`
	MOMID           string `json:"mom_id,omitempty" form:"mom_id" xml:"mom_id"`
	IncludeChildren bool   `json:"include_children" form:"include_children" xml:"include_children"`
}

// NotificationRemoveTokenRequest is a type holding a request to remove a token from a notification
type NotificationRemoveTokenRequest struct {
	Mytoken string `json:"mytoken,omitempty" form:"mytoken" xml:"mytoken"`
	MOMID   string `json:"mom_id,omitempty" form:"mom_id" xml:"mom_id"`
}

// NotificationUpdateNotificationClassesRequest is a type holding a request to update the notification classes of a
// notification
type NotificationUpdateNotificationClassesRequest struct {
	Classes []*NotificationClass `json:"notification_classes"`
}

// NotificationsListResponse is a type holding the response to a notification list request
type NotificationsListResponse struct {
	Notifications []NotificationInfo `json:"notifications"`
	TokenUpdate   *MytokenResponse   `json:"token_update,omitempty"`
}

// NotificationClass is a type for notification classes
type NotificationClass struct {
	Name           string
	Description    string
	relevantEvents []Event
	children       []*NotificationClass
}

// Defined NotificationClasses
var (
	NotificationClassATs = &NotificationClass{
		Name:           "AT_creations",
		Description:    "Notifications for the creation of access tokens",
		relevantEvents: []Event{EventATCreated},
	}
	NotificationClassMTs = &NotificationClass{
		Name:           "subtoken_creations",
		Description:    "Notifications for the creation of mytoken subtokens",
		relevantEvents: []Event{EventMTCreated},
	}
	NotificationClassSettingChanges = &NotificationClass{
		Name:        "setting_changes",
		Description: "Notifications for changes in the settings",
		relevantEvents: []Event{
			EventCalendarCreated,
			EventCalendarDeleted,
			EventEmailChanged,
			EventEmailMimetypeChanged,
			EventSSHKeyAdded,
			EventGrantEnabled,
			EventGrantDisabled,
		},
	}
	NotificationClassSecurity = &NotificationClass{
		Name:        "security",
		Description: "Notifications for all security related events",
	}
	NotificationClassBlockedUsages = &NotificationClass{
		Name:        subNotificationClassName(NotificationClassSecurity, "blocked_usages"),
		Description: "Notifications for blocked usages",
	}
	NotificationClassInsufficientCapabilities = &NotificationClass{
		Name:        subNotificationClassName(NotificationClassBlockedUsages, "capabilities"),
		Description: "Notifications for blocked usages because of insufficient capabilities",
	}
	NotificationClassRestrictedUsages = &NotificationClass{
		Name:        subNotificationClassName(NotificationClassBlockedUsages, "restrictions"),
		Description: "Notifications for blocked usages because of restrictions",
	}
	NotificationClassRevokedUsage = &NotificationClass{
		Name:        subNotificationClassName(NotificationClassSecurity, "revoked"),
		Description: "Notifications for tried usage of a revoked token (only works as a user-wide notification)",
	}
	NotificationClassUnusualIPs = &NotificationClass{
		Name:        subNotificationClassName(NotificationClassSecurity, "ips"),
		Description: "Notifications for usages from previously unknown ip addresses",
	}
)

func init() {
	NotificationClassSecurity.children = []*NotificationClass{
		NotificationClassUnusualIPs,
		// NotificationsClassUnusualCountries,
		NotificationClassBlockedUsages,
		NotificationClassRevokedUsage,
	}
	NotificationClassBlockedUsages.children = []*NotificationClass{
		NotificationClassInsufficientCapabilities,
		NotificationClassRestrictedUsages,
	}
}

// AllNotificationClasses holds all defined NotificationClasses
var AllNotificationClasses = []*NotificationClass{
	NotificationClassATs,
	NotificationClassMTs,
	NotificationClassSettingChanges,
	NotificationClassSecurity,
	NotificationClassBlockedUsages,
	NotificationClassInsufficientCapabilities,
	NotificationClassRestrictedUsages,
}

func subNotificationClassName(nc *NotificationClass, Suffix string) string {
	return nc.Name + ":" + Suffix
}

// NewNotificationClass casts a string into a NotificationClass
func NewNotificationClass(name string) *NotificationClass {
	for _, nc := range AllNotificationClasses {
		if nc.Name == name {
			return nc
		}
	}
	return &NotificationClass{Name: name}
}

// MarshalJSON implements the json.Marshaler interface
func (nc NotificationClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(nc.Name)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (nc *NotificationClass) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*nc = *NewNotificationClass(name)
	return nil
}

// Contains checks if the NotificationClass contains another NotificationClass (i.e. as a subclass)
func (nc NotificationClass) Contains(v *NotificationClass) bool {
	if nc.Name == v.Name {
		return true
	}
	for _, nn := range nc.children {
		if nn.Contains(v) {
			return true
		}
	}
	return false
}

// GetChildren returns the children of a NotificationClass
func (nc *NotificationClass) GetChildren() []*NotificationClass {
	return nc.children
}

var notificationClassEventMap map[string]*NotificationClass

func init() {
	notificationClassEventMap = make(map[string]*NotificationClass)
	for _, nc := range AllNotificationClasses {
		for _, e := range nc.relevantEvents {
			notificationClassEventMap[e.String()] = nc
		}
	}
}

// NotificationClassFromEvent returns the *NotificationClass linked to an Event
func NotificationClassFromEvent(event Event) *NotificationClass {
	nc, ok := notificationClassEventMap[event.String()]
	if ok {
		return nc
	}
	return nil
}
