package domain

import "time"

// TroubleTicket represents the TMF621 entity for managing incidents
type TroubleTicket struct {
	ID             string    \`json:"id" gorm:"primaryKey"\`
	ExternalID     string    \`json:"externalId,omitempty"\`
	Description    string    \`json:"description"\`
	Severity       string    \`json:"severity"\` // Critical, Major, Minor
	Status         string    \`json:"status"\`   // Open, InProgress, Resolved, Closed
	CreationDate   time.Time \`json:"creationDate"\`
	TargetResolutionDate time.Time \`json:"targetResolutionDate,omitempty"\`
	RelatedParty   []RelatedParty \`json:"relatedParty" gorm:"serializer:json"\`
	Note           []Note         \`json:"note" gorm:"serializer:json"\`
}

type RelatedParty struct {
	ID   string \`json:"id"\`
	Role string \`json:"role"\` // Customer, Technical Lead
}

type Note struct {
	Author string    \`json:"author"\`
	Date   time.Time \`json:"date"\`
	Text   string    \`json:"text"\`
}
