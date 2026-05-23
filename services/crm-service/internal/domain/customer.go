package domain

import (
	"time"
)

// Customer represents the TMF629 Customer entity
type Customer struct {
	ID             string          `json:"id" gorm:"primaryKey"`
	Name           string          `json:"name"`
	Status         string          `json:"status"` // Active, Suspended, Terminated
	CustomerType   string          `json:"customerType"` // Individual, Organization
	ContactMedium  []ContactMedium `json:"contactMedium" gorm:"serializer:json"`
	CreditProfile  CreditProfile   `json:"creditProfile" gorm:"serializer:json"`
	Accounts       []AccountRef    `json:"accounts" gorm:"serializer:json"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

type ContactMedium struct {
	MediumType string `json:"mediumType"` // Email, Phone, Social
	Preferred  bool   `json:"preferred"`
	Value      string `json:"value"`
}

type CreditProfile struct {
	CreditScore int    `json:"creditScore"`
	CreditLimit float64 `json:"creditLimit"`
}

type AccountRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}
