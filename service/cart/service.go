package cart

import (
	"fmt"

	"github.com/saikumaradapa/ecom/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductId)
		}

		productIds[i] = item.ProductId
	}

	return productIds, nil
}

func CreateOrder(ps []types.Product, item []types.CartItem, userID int) (int, float64, error) {
	// to do
	return 0, 0, nil
}
