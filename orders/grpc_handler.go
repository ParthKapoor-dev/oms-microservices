package main

import (
	"context"
	"log"

	pb "github.com/parthkapoor-dev/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGrpcHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {

	log.Printf("New Order Received, %v", p)
	o := &pb.Order{
		ID: "45",
	}
	return o, nil
}
