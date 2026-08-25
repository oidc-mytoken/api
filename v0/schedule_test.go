package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScheduleValidate(t *testing.T) {
	tests := []struct {
		name    string
		sched   Schedule
		wantErr bool
	}{
		{"empty is valid", Schedule{}, false},
		{"timezone valid", Schedule{Timezone: "Europe/Berlin"}, false},
		{"timezone invalid", Schedule{Timezone: "Not/AZone"}, true},
		{"days valid", Schedule{Days: []int{0, 1, 7}}, false},
		{"days out of range", Schedule{Days: []int{8}}, true},
		{"days negative", Schedule{Days: []int{-1}}, true},
		{"dom valid", Schedule{DaysOfMonth: []int{1, 15, -1, -31}}, false},
		{"dom zero", Schedule{DaysOfMonth: []int{0}}, true},
		{"dom out of range", Schedule{DaysOfMonth: []int{32}}, true},
		{"window valid", Schedule{From: "09:00", To: "17:00"}, false},
		{"window wrap valid", Schedule{From: "22:00", To: "02:00"}, false},
		{"window only from", Schedule{From: "09:00"}, true},
		{"window only to", Schedule{To: "17:00"}, true},
		{"window bad format", Schedule{From: "9:00", To: "17:00"}, true},
		{"window out of range", Schedule{From: "24:00", To: "17:00"}, true},
		{"every valid", Schedule{Every: 2}, false},
		{"every negative", Schedule{Every: -1}, true},
		{"anchor valid", Schedule{Every: 2, Anchor: "2026-01-01"}, false},
		{"anchor bad format", Schedule{Anchor: "01/01/2026"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.sched.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestScheduleUnmarshalJSON(t *testing.T) {
	var s Schedule
	err := json.Unmarshal([]byte(`{"timezone":"UTC","days":[1,5],"from":"09:00","to":"17:00"}`), &s)
	require.NoError(t, err)
	assert.Equal(t, Schedule{
		Timezone: "UTC",
		Days:     []int{1, 5},
		From:     "09:00",
		To:       "17:00",
	}, s)

	var bad Schedule
	err = json.Unmarshal([]byte(`{"days":[9]}`), &bad)
	assert.Error(t, err)

	var badTZ Schedule
	err = json.Unmarshal([]byte(`{"timezone":"Mars/Olympus"}`), &badTZ)
	assert.Error(t, err)
}

func TestRestrictionSchedule(t *testing.T) {
	sched := &Schedule{Timezone: "UTC", Days: []int{1}, From: "09:00", To: "17:00"}
	r := Restriction{Schedule: sched}
	b, err := json.Marshal(r)
	require.NoError(t, err)
	var back Restriction
	require.NoError(t, json.Unmarshal(b, &back))
	require.NotNil(t, back.Schedule)
	assert.Equal(t, sched, back.Schedule)
}

func TestRestrictionClaimSchedule(t *testing.T) {
	assert.Equal(t, "schedule", RestrictionClaimSchedule)
	assert.Contains(t, AllRestrictionClaims, RestrictionClaimSchedule)
}
