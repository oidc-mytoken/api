package api

import (
	"database/sql/driver"

	"github.com/pkg/errors"
)

// Event is a type for events
type Event string

func (e Event) String() string {
	return string(e)
}

// Value implements the sql.Valuer interface
func (e Event) Value() (driver.Value, error) {
	return e.String(), nil
}

// Scan implements the sql.Scanner interface
func (e *Event) Scan(src interface{}) error {
	var eventString string
	switch src := src.(type) {
	case string:
		eventString = src
	case []byte:
		eventString = string(src)
	default:
		return errors.New("incompatible type for Event")
	}
	*e = Event(eventString)
	return nil
}

// NewEvent creates a new Event from the event string
func NewEvent(typ string) Event {
	return Event(typ)
}

// Events for Mytokens
var (
	EventUnknown                          = Event("unknown")
	EventMTCreated                        = Event("created")
	EventATCreated                        = Event("AT_created")
	EventSubtokenCreated                  = Event("MT_created")
	EventTokenInfoIntrospect              = Event("tokeninfo_introspect")
	EventTokenInfoHistory                 = Event("tokeninfo_history")
	EventTokenInfoSubtokens               = Event("tokeninfo_subtokens")
	EventTokenInfoListMTs                 = Event("tokeninfo_list_mytokens")
	EventTokenInfoNotifications           = Event("tokeninfo_notifications")
	EventInheritedRT                      = Event("inherited_RT")
	EventTransferCodeCreated              = Event("transfer_code_created")
	EventTransferCodeUsed                 = Event("transfer_code_used")
	EventMTRotated                        = Event("token_rotated")
	EventGrantEnabled                     = Event("settings_grant_enabled")
	EventGrantDisabled                    = Event("settings_grant_disabled")
	EventGrantsListed                     = Event("settings_grants_listed")
	EventSSHKeyListed                     = Event("ssh_keys_listed")
	EventSSHKeyAdded                      = Event("ssh_key_added")
	EventRevokedOtherToken                = Event("revoked_other_token")
	EventTokenInfoHistoryOtherToken       = Event("tokeninfo_history_other_token")
	EventTokenInfoNotificationsOtherToken = Event("tokeninfo_notifications_other_token")
	EventExpired                          = Event("expired")
	EventRevoked                          = Event("revoked")
	EventNotificationSubscribed           = Event("notification_subscribed")
	EventNotificationListed               = Event("notification_listed")
	EventNotificationUnsubscribed         = Event("notification_unsubscribed")
	EventNotificationSubscribedOther      = Event("notification_subscribed_other")
	EventNotificationUnsubscribedOther    = Event("notification_unsubscribed_other")
	EventNotificationCreated              = Event("notification_created")
	EventNotificationCreatedOther         = Event("notification_created_other")
	EventCalendarCreated                  = Event("calendar_created")
	EventCalendarListed                   = Event("calendar_listed")
	EventCalendarDeleted                  = Event("calendar_deleted")
	EventEmailSettingsListed              = Event("email_settings_listed")
	EventEmailChanged                     = Event("email_changed")
	EventEmailMimetypeChanged             = Event("email_mimetype_changed")
)
