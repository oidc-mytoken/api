package api

// EventHistory is a slice of EventEntry
type EventHistory struct {
	Events []EventEntry `json:"events"`
}

// EventEntry is a type holding information about an event
type EventEntry struct {
	Event          string `db:"event" json:"event"`
	Time           int64  `db:"time" json:"time"`
	Comment        string `db:"comment" json:"comment,omitempty"`
	ClientMetaData `json:",inline"`
	MOMID          string `json:"mom_id,omitempty"`
}
