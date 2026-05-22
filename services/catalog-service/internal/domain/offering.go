package domain

import (
	"time"
)

// ProductOffering represents the TMF620 Product Offering entity
type ProductOffering struct {
	ID               string            \`json:"id" gorm:"primaryKey"\`
	Name             string            \`json:"name"\`
	Description      string            \`json:"description"\`
	Status           string            \`json:"status"\` // Active, Retired
	Version          string            \`json:"version"\`
	LifecycleStatus  string            \`json:"lifecycleStatus"\`
	IsBundle         bool              \`json:"isBundle"\`
	ProductPrices    []ProductPrice    \`json:"productOfferingPrice" gorm:"serializer:json"\`
	ProductSpec      ProductSpecRef    \`json:"productSpecification" gorm:"serializer:json"\`
	BundleComponents []ComponentRef    \`json:"bundledProductOffering,omitempty" gorm:"serializer:json"\`
	ValidFor         TimePeriod        \`json:"validFor" gorm:"serializer:json"\`
}

type ProductPrice struct {
	ID          string     \`json:"id"\`
	Name        string     \`json:"name"\`
	PriceType   string     \`json:"priceType"\` // Recurring, One-Time, Usage
	Price       PriceValue \`json:"price"\`
	Period      string     \`json:"recurringChargePeriod,omitempty"\`
}

type PriceValue struct {
	Amount   float64 \`json:"amount"\`
	Currency string  \`json:"currency"\`
}

type ProductSpecRef struct {
	ID   string \`json:"id"\`
	Name string \`json:"name"\`
}

type ComponentRef struct {
	ID   string \`json:"id"\`
	Name string \`json:"name"\`
}

type TimePeriod struct {
	StartDateTime time.Time \`json:"startDateTime"\`
	EndDateTime   time.Time \`json:"endDateTime"\`
}
