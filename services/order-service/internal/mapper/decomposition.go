package mapper

import (
	"fmt"
	"github.com/telcoflow/telcoflow/services/order-service/internal/domain"
)

// DecomposeProductOrder maps a Product Order (Customer perspective) to Service Orders (Network perspective)
func DecomposeProductOrder(productOrder domain.ProductOrder) ([]string, error) {
	var serviceOrders []string

	for _, item := range productOrder.OrderItems {
		// Example: A "Fiber Broadband" product offering decomposes into:
		// 1. PPPoE Service Activation
		// 2. ONT Configuration
		// 3. Billing Activation

		switch item.Offering.Name {
		case "Fiber Broadband":
			serviceOrders = append(serviceOrders,
				fmt.Sprintf("SO-PPPOE-%s", item.ID),
				fmt.Sprintf("SO-ONT-%s", item.ID),
			)
		case "Mobile 5G":
			serviceOrders = append(serviceOrders,
				fmt.Sprintf("SO-HSS-%s", item.ID),
			)
		default:
			serviceOrders = append(serviceOrders,
				fmt.Sprintf("SO-GENERIC-%s", item.ID),
			)
		}
	}

	return serviceOrders, nil
}
