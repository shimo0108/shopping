package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb_food "github.com/shimo0108/shopping/backend/server/pb/food"
	pb_line_food "github.com/shimo0108/shopping/backend/server/pb/line_food"
	pb_order "github.com/shimo0108/shopping/backend/server/pb/order"
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

	foodPersistence := persistence.NewFoodPersistence(db.NewDB())
	foodUsecase := usecase.NewFoodUsecase(foodPersistence)
	foodHandler := handler.NewFoodHandler(foodUsecase)

	lineFoodPersistence := persistence.NewLineFoodPersistence(db.NewDB())
	lineFoodUsecase := usecase.NewLineFoodUsecase(lineFoodPersistence)
	lineFoodHandler := handler.NewLineFoodHandler(lineFoodUsecase)

	orderPersistence := persistence.NewOrderPersistence(db.NewDB())
	orderUsecase := usecase.NewOrderUsecase(orderPersistence)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	fmt.Println("Go gRPC START!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb_test.RegisterTestServiceServer(grpcServer, &TestService{})
	pb_restaurant.RegisterRestaurantServiceServer(grpcServer, restaurantHandler)
	pb_food.RegisterFoodServiceServer(grpcServer, foodHandler)
	pb_line_food.RegisterLineFoodServiceServer(grpcServer, lineFoodHandler)
	pb_order.RegisterOrderServiceServer(grpcServer, orderHandler)

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
