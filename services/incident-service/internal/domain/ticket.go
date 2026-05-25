package domain

import "time"

type TroubleTicket struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	Description     string    `json:"description"`
	Severity        string    `json:"severity"`
	Status          string    `json:"status"` // Open, InProgress, Resolved, Closed
	CreationDate    time.Time `json:"creationDate"`
	ResolutionDate  time.Time `json:"resolutionDate,omitempty"`
	CustomerID      string    `json:"customerId"`
	RelatedObjectID string    `json:"relatedObjectId"`
}
