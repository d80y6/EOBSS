package domain

import "fmt"

// ValidateBundle checks if all component requirements are met for a bundled product offering
func (p *ProductOffering) ValidateBundle(availableComponents []string) error {
	if !p.IsBundle {
		return nil
	}

	for _, req := range p.BundleComponents {
		found := false
		for _, available := range availableComponents {
			if req.ID == available {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("bundle validation failed: missing component %s (%s)", req.Name, req.ID)
		}
	}

	return nil
}

// CalculateDynamicPrice applies context-aware discounts to the offering price
func (p *ProductOffering) CalculateDynamicPrice(customerTier string) []ProductPrice {
	updatedPrices := make([]ProductPrice, len(p.ProductPrices))
	copy(updatedPrices, p.ProductPrices)

	if customerTier == "Gold" {
		for i := range updatedPrices {
			// Apply a 10% loyalty discount for Gold customers
			updatedPrices[i].Price.Amount *= 0.9
		}
	}

	return updatedPrices
}
