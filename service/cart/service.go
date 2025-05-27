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

func (h *Handler) CreateOrder(ps []types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}

	// check if all products are actually in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	// calculate the total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce the quantity of products in our db
	for _, item := range items {
		product := productMap[item.ProductId]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

	// create the order
	orderId, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "some address",
	})
	if err != nil {
		return 0, 0, err
	}

	// create order items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderId,
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductId].Price,
		})
	}

	return orderId, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductId]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductId)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %d is not available in the quantity requested", item.Quantity)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductId]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
