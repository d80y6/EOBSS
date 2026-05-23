package domain

import (
	"time"
)

// Alarm represents the TMF642 Alarm entity
type Alarm struct {
	ID                 string    `json:"id" gorm:"primaryKey"`
	ExternalID         string    `json:"externalId"`
	AlarmType          string    `json:"alarmType"`
	Severity           string    `json:"severity"` // Critical, Major, Minor, Warning, Indeterminate
	ProbableCause      string    `json:"probableCause"`
	SpecificProblem    string    `json:"specificProblem"`
	AlarmedResource    ResourceRef `json:"alarmedResource" gorm:"serializer:json"`
	AlarmRaisedTime    time.Time `json:"alarmRaisedTime"`
	AlarmClearedTime   *time.Time `json:"alarmClearedTime,omitempty"`
	AckStatus          string    `json:"ackStatus"` // unacknowledged, acknowledged
	State              string    `json:"state"` // active, cleared
}

type ResourceRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
