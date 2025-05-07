package main

import (
	"context"
	"log"

	pb "github.com/parthkapoor-dev/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrderService
}

func NewGrpcHandler(grpcServer *grpc.Server, service OrderService) {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {

	log.Printf("New Order Received, Order: %v", p)

	if err := h.service.ValidateOrder(context.Background(), p); err != nil {
		return nil, err
	}

	o := &pb.Order{
		ID: "45",
	}
	return o, nil
}
