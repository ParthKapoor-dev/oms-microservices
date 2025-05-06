package main

import (
	"context"
	"net/http"

	"github.com/parthkapoor-dev/common"
	pb "github.com/parthkapoor-dev/common/api"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders/", h.handleCreateOrder)
}

func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.client.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
}
