package main

import (
	"net/http"

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
}
