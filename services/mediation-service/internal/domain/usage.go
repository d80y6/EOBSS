package domain

import "time"

type UsageEvent struct {
	ID        string    `json:"id"`
	UsageType string    `json:"usageType"`
	Quantity  float64   `json:"quantity"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
	ServiceID string    `json:"serviceId"`
}
