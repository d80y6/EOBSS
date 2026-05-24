package domain

import "time"

type Invoice struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customerId"`
	Amount     float64   `json:"amount"`
	Currency   string    `json:"currency"`
	DueDate    time.Time `json:"dueDate"`
	Status     string    `json:"status"`
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
