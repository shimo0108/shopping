package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb_restaurant "github.com/shimo0108/shopping/backend/server/pb/restaurant"
	pb_test "github.com/shimo0108/shopping/backend/server/pb/test"
	"github.com/shimo0108/shopping/backend/server/shopping/config/db"
	_ "github.com/shimo0108/shopping/backend/server/shopping/config/db"
	"github.com/shimo0108/shopping/backend/server/shopping/infrastructure/persistence"
	"github.com/shimo0108/shopping/backend/server/shopping/interface/handler"
	"github.com/shimo0108/shopping/backend/server/shopping/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type TestService struct {
	pb_test.UnimplementedTestServiceServer
}

func main() {

	restaurantPersistence := persistence.NewRestaurantPersistence(db.NewDB())
	restaurantUsecase := usecase.NewRestaurantUsecase(restaurantPersistence)
	restaurantHandler := handler.NewRestaurantHandler(restaurantUsecase)

	fmt.Println("Go gRPC START!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb_test.RegisterTestServiceServer(grpcServer, &TestService{})
	pb_restaurant.RegisterRestaurantServiceServer(grpcServer, restaurantHandler)
	log.Printf("gRPC server start")

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (grpcServer *TestService) Test(ctx context.Context, message *pb_test.TestRequest) (*pb_test.TestResponse, error) {
	log.Printf("Received: %v", message.Message)

	return &pb_test.TestResponse{Message: "Test " + message.Message}, nil
}
