package api

type MailSettingsInfoResponse struct {
	EmailAddress   string `json:"email_address"`
	EmailVerified  bool   `json:"email_verified"`
	PreferHTMLMail bool   `json:"prefer_html_mail"`
}

type UpdateMailSettingsRequest struct {
	EmailAddress   string `json:"email_address,omitempty"`
	PreferHTMLMail *bool  `json:"prefer_html_mail,omitempty"`
}
