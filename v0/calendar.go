package api

type NotificationCalendar struct {
	Name    string `json:"name"`
	ICSPath string `json:"ics_path"`
}

type AddMytokenToCalendarRequest struct {
	MomID   string `json:"mom_id" xml:"mom_id" form:"mom_id"`
	Comment string `json:"comment" xml:"comment" form:"comment"`
}
