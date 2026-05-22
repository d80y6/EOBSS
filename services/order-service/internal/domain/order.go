package domain
type ProductOrder struct {
	ID    string `json:"id" gorm:"primaryKey"`
	State string `json:"state"`
}
