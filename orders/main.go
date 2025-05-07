package main

import (
	"log"
	"net"

	"github.com/parthkapoor-dev/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen at %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGrpcHandler(grpcServer, svc)

	log.Println("grpc server started at ", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
