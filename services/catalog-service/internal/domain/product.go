package domain
import "time"
type ProductOffering struct {
	ID             string            `json:"id" gorm:"primaryKey"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	Status         string            `json:"status"`
}
