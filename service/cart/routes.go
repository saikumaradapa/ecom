package cart

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saikumaradapa/ecom/types"
	"github.com/saikumaradapa/ecom/utils"
)

type Handler struct {
	store types.OrderStore
	productStore types.ProductStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request)  {
	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return 
	}

	// get products 
}