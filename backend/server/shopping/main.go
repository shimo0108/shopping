package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/shimo0108/shopping/backend/server/pb/test"

	"google.golang.org/grpc"
)

type TestService struct {
	pb.UnimplementedTestServiceServer
}

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTestServiceServer(grpcServer, &TestService{})
	log.Printf("gRPC server start")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (grpcServer *TestService) Test(ctx context.Context, message *pb.TestRequest) (*pb.TestResponse, error) {
	log.Printf("Received: %v", message.Message)

	return &pb.TestResponse{Message: "Test " + message.Message}, nil
}
