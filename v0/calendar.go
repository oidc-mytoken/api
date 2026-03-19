package api

type CreateCalendarRequest struct {
	Description string `json:"description,omitempty"`
	Tags        []Tag  `json:"tags"`
}

type NotificationCalendar struct {
	ICSPath     string    `json:"ics_path"`
	Description string    `json:"description,omitempty"`
	Tags        []TagInfo `json:"tags"`
}

type CalendarInfo struct {
	NotificationCalendar
	SubscribedTokens []string `json:"subscribed_tokens,omitempty"`
}

type AddMytokenToCalendarRequest struct {
	MomID   string `json:"mom_id" xml:"mom_id" form:"mom_id"`
	Comment string `json:"comment" xml:"comment" form:"comment"`
}

// CalendarListResponse is the response returned to list all calendars of a user
type CalendarListResponse struct {
	Calendars   []CalendarInfo   `json:"calendars"`
	TokenUpdate *MytokenResponse `json:"token_update,omitempty"`
}
