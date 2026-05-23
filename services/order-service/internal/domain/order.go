package domain

import (
	"time"
)

// ProductOrder represents the TMF622 Product Order entity
type ProductOrder struct {
	ID                   string              `json:"id" gorm:"primaryKey"`
	ExternalID           string              `json:"externalId"`
	OrderDate            time.Time           `json:"orderDate"`
	State                string              `json:"state"` // Acknowledged, InProgress, Completed, Failed
	OrderType            string              `json:"orderType"` // New, Change, Termination
	RequestedStartDate   time.Time           `json:"requestedStartDate"`
	RequestedCompletionDate time.Time        `json:"requestedCompletionDate"`
	Customer             CustomerRef         `json:"customer" gorm:"serializer:json"`
	OrderItems           []OrderItem         `json:"orderItem" gorm:"serializer:json"`
	RelatedParties       []RelatedParty      `json:"relatedParty" gorm:"serializer:json"`
	Notes                []Note              `json:"note" gorm:"serializer:json"`
}

type OrderItem struct {
	ID           string           `json:"id"`
	Action       string           `json:"action"` // Add, Change, Remove
	State        string           `json:"state"`
	Product      ProductRef       `json:"product"`
	Offering     OfferingRef      `json:"productOffering"`
}

type CustomerRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductRef struct {
	ID string `json:"id"`
}

type OfferingRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RelatedParty struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type Note struct {
	Author string    `json:"author"`
	Date   time.Time `json:"date"`
	Text   string    `json:"text"`
}
