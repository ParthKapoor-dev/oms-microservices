package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/parthkapoor-dev/common"
	pb "github.com/parthkapoor-dev/common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders/", h.handleCreateOrder)
}

func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
	}

	o, err := h.client.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	if rStatus := status.Convert(err); rStatus != nil {

		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}

		common.WriteError(w, http.StatusInternalServerError, err.Error())
	}

	common.WriteJSON(w, http.StatusOK, o)
}

func validateItems(items []*pb.ItemsWithQuantity) error {

	if len(items) == 0 {
		return errors.New("Should have atleast 1 Item")
	}

	for _, i := range items {
		if i.ID == "" {
			return errors.New("All Items should have some Id")
		}

		if i.Quantity <= 0 {
			return errors.New("Quantity should be greater than 0")
		}
	}

	return nil

}
