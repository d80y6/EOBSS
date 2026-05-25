package domain

import "time"

type Invoice struct {
	ID         string        `json:"id"`
	CustomerID string        `json:"customerId"`
	Amount     Money         `json:"amount"`
	DueDate    time.Time     `json:"dueDate"`
	Status     string        `json:"status"`
	Items      []InvoiceItem `json:"invoiceItem"`
}

type InvoiceItem struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Amount      Money   `json:"amount"`
	ServiceID   string  `json:"serviceId"`
}

type Money struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type UsageRecord struct {
	ID            string    `json:"id"`
	ServiceID     string    `json:"serviceId"`
	UsageType     string    `json:"usageType"`
	UsageQuantity float64   `json:"usageQuantity"`
	UsageUnit     string    `json:"usageUnit"`
	RatedAmount   float64   `json:"ratedAmount"`
	Timestamp     time.Time `json:"timestamp"`
}
