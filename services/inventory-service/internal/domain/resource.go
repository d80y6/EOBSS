package domain

import (
	"time"
)

// Resource represents the TMF639 Resource entity
type Resource struct {
	ID                 string             \`json:"id" gorm:"primaryKey"\`
	Name               string             \`json:"name"\`
	Description        string             \`json:"description"\`
	ResourceStatus     string             \`json:"resourceStatus"\` // available, reserved, inactive, active
	UsageState         string             \`json:"usageState"\`     // idle, active, busy
	ResourceSpec       ResourceSpecRef    \`json:"resourceSpecification" gorm:"serializer:json"\`
	Characteristics    []Characteristic   \`json:"resourceCharacteristic" gorm:"serializer:json"\`
	RelatedResources   []RelatedResource  \`json:"relatedResource" gorm:"serializer:json"\`
	CreatedAt          time.Time          \`json:"createdAt"\`
	UpdatedAt          time.Time          \`json:"updatedAt"\`
}

type ResourceSpecRef struct {
	ID   string \`json:"id"\`
	Name string \`json:"name"\`
}

type Characteristic struct {
	Name  string \`json:"name"\`
	Value string \`json:"value"\`
}

type RelatedResource struct {
	ID   string \`json:"id"\`
	Role string \`json:"role"\`
}
