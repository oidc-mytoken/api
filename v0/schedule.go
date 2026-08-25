package api

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Schedule describes a periodic time window in which a mytoken may be used.
//
// All fields are optional. A field that is not set does not restrict the
// corresponding dimension:
//   - timezone: IANA timezone name; defaults to UTC
//   - days: weekdays on which usage is allowed, 1 (Monday) to 7 (Sunday);
//     0 is accepted as an alias for 7 (Sunday)
//   - days_of_month: days of month on which usage is allowed, 1 to 31;
//     negative values count from the end of the month, e.g. -1 is the last day
//   - from/to: allowed time window of a day in 24h HH:MM format; to may be
//     smaller than from to describe a window wrapping around midnight;
//     from == to describes a full day
//   - every: allow usage only every N days starting from anchor
//   - anchor: reference date for every in YYYY-MM-DD format; defaults to the
//     creation date of the mytoken
type Schedule struct {
	Timezone    string `json:"timezone,omitempty"`
	Days        []int  `json:"days,omitempty"`
	DaysOfMonth []int  `json:"days_of_month,omitempty"`
	From        string `json:"from,omitempty"`
	To          string `json:"to,omitempty"`
	Every       int    `json:"every,omitempty"`
	Anchor      string `json:"anchor,omitempty"`
}

// Validate checks that the schedule is well-formed.
func (s *Schedule) Validate() error {
	if s.Timezone != "" {
		if _, err := time.LoadLocation(s.Timezone); err != nil {
			return errors.Errorf("invalid timezone '%s' in schedule", s.Timezone)
		}
	}
	for _, d := range s.Days {
		if d < 0 || d > 7 {
			return errors.Errorf(
				"invalid day of week %d in schedule; allowed values are 0 to 7 where 0 and 7 are Sunday", d,
			)
		}
	}
	for _, d := range s.DaysOfMonth {
		if d == 0 || d < -31 || d > 31 {
			return errors.Errorf(
				"invalid day of month %d in schedule; allowed values are -31 to 31 excluding 0", d,
			)
		}
	}
	if s.From != "" || s.To != "" {
		if s.From == "" || s.To == "" {
			return errors.New("from and to must either both be set or both be empty in schedule")
		}
		for _, t := range []string{s.From, s.To} {
			tt, err := time.Parse("15:04", t)
			if err != nil || tt.Format("15:04") != t {
				return errors.Errorf("invalid time '%s' in schedule; expected format HH:MM", t)
			}
		}
	}
	if s.Every < 0 {
		return errors.Errorf("invalid every value %d in schedule; must be positive", s.Every)
	}
	if s.Anchor != "" {
		if _, err := time.Parse("2006-01-02", s.Anchor); err != nil {
			return errors.Errorf("invalid anchor '%s' in schedule; expected format YYYY-MM-DD", s.Anchor)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *Schedule) UnmarshalJSON(data []byte) error {
	type alias Schedule
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return errors.WithStack(err)
	}
	*s = Schedule(tmp)
	return errors.WithStack(s.Validate())
}
