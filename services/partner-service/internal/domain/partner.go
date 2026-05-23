package domain

import "time"

// Partner represents the TMF680 Partner entity for wholesale/interconnect
type Partner struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	PartnerType  string    `json:"partnerType"` // MVNO, Interconnect, Reseller
	Status       string    `json:"status"`
	ValidFor     TimePeriod `json:"validFor" gorm:"serializer:json"`
	Account      AccountRef `json:"account" gorm:"serializer:json"`
}

type TimePeriod struct {
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime,omitempty"`
}

type AccountRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
