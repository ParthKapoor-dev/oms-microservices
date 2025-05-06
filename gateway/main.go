package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	common "github.com/parthkapoor-dev/common"
	pb "github.com/parthkapoor-dev/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	OrderServiceAddr = "localhost:3000"
)

func main() {

	conn, err := grpc.NewClient(OrderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to Dial up the server, %v", err)
	}
	defer conn.Close()

	log.Println("Dialing Order Service at ", OrderServiceAddr)

	orderClient := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(orderClient)
	handler.registerRoutes(mux)

	log.Printf("Starting the server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to Start the server")
	}

}
