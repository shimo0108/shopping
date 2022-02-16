package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw_food "github.com/shimo0108/shopping/backend/server/pb/food"
	gw_restaurant "github.com/shimo0108/shopping/backend/server/pb/restaurant"
	gw_test "github.com/shimo0108/shopping/backend/server/pb/test"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	testEndpoint       = "shop_grpc:8080"
	restaurantEndpoint = "shop_grpc:8080"
	foodEndpoint       = "shop_grpc:8080"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	test_err := gw_test.RegisterTestServiceHandlerFromEndpoint(ctx, mux, testEndpoint, opts)
	if test_err != nil {
		fmt.Printf("serve: %v\n", test_err)
	}
	restaurant_err := gw_restaurant.RegisterRestaurantServiceHandlerFromEndpoint(ctx, mux, restaurantEndpoint, opts)
	if restaurant_err != nil {
		fmt.Printf("serve: %v\n", restaurant_err)
	}
	food_err := gw_food.RegisterFoodServiceHandlerFromEndpoint(ctx, mux, foodEndpoint, opts)
	if food_err != nil {
		fmt.Printf("serve: %v\n", food_err)
	}

	return http.ListenAndServe(":9999", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
