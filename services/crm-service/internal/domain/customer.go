package domain
type Customer struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
