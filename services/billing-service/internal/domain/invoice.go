package domain

import (
	"time"
)

// Invoice represents the TMF678 Customer Bill entity
type Invoice struct {
	ID              string          \`json:"id" gorm:"primaryKey"\`
	BillingAccount  AccountRef      \`json:"billingAccount" gorm:"serializer:json"\`
	InvoiceDate     time.Time       \`json:"invoiceDate"\`
	DueDate         time.Time       \`json:"dueDate"\`
	TotalAmount     Money           \`json:"totalAmount" gorm:"serializer:json"\`
	TaxAmount       Money           \`json:"taxAmount" gorm:"serializer:json"\`
	RemainingAmount Money           \`json:"remainingAmount" gorm:"serializer:json"\`
	Status          string          \`json:"status"\` // Draft, Issued, Paid, Overdue
	Items           []InvoiceItem   \`json:"invoiceItem" gorm:"serializer:json"\`
}

type InvoiceItem struct {
	ID          string    \`json:"id"\`
	Description string    \`json:"description"\`
	Quantity    float64   \`json:"quantity"\`
	UnitPrice   Money     \`json:"unitPrice"\`
	TotalAmount Money     \`json:"totalAmount"\`
	ServiceRef  string    \`json:"serviceRef,omitempty"\`
}

type AccountRef struct {
	ID   string \`json:"id"\`
	Name string \`json:"name"\`
}

type Money struct {
	Amount   float64 \`json:"amount"\`
	Currency string  \`json:"currency"\`
}

// UsageRecord (CDR) represents a mediated usage event
type UsageRecord struct {
	ID            string    \`json:"id"\`
	UsageType     string    \`json:"usageType"\` // Data, Voice, SMS
	UsageQuantity float64   \`json:"usageQuantity"\`
	UsageUnit     string    \`json:"usageUnit"\` // MB, Minutes, Count
	Timestamp     time.Time \`json:"timestamp"\`
	ServiceRef    string    \`json:"serviceRef"\`
	RatedAmount   float64   \`json:"ratedAmount,omitempty"\`
}
