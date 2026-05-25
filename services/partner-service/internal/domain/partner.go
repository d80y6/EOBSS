package domain

import (
	"time"
)

// Partner represents TMF680 Partner entity for settlements
type Partner struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Type         string    `json:"type"` // MVNO, Wholesaler, ContentProvider
	Status       string    `json:"status"`
	OnboardDate  time.Time `json:"onboardDate"`
	SettlementID string    `json:"settlementId"`
}
