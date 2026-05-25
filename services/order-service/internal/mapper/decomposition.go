package mapper

import (
	"fmt"
	"github.com/telcoflow/telcoflow/services/order-service/internal/domain"
)

// OrderDecomposer maps TMF622 Product Orders to TMF641 Service Orders
type OrderDecomposer struct{}

func (d *OrderDecomposer) Decompose(order domain.ProductOrder) ([]string, error) {
	var serviceOrders []string

	for _, item := range order.OrderItems {
		// Logic: If Product is 'Broadband', decompose to 'Radius' and 'Port' services
		// In a real system, this would query the Catalog for Product-to-Service mappings

		serviceType := "Generic"
		if item.Product.Name == "Broadband-Fiber" {
			serviceType = "Radius"
		} else if item.Product.Name == "VoIP-Mobile" {
			serviceType = "Kamailio"
		}

		soID := fmt.Sprintf("SO-%s-%s", order.ID, serviceType)
		serviceOrders = append(serviceOrders, soID)
	}

	return serviceOrders, nil
}
