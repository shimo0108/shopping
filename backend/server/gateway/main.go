package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pw_test "github.com/shimo0108/shopping/backend/server/pb/test"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("shop_grpc:8080")
	err := pw_test.RegisterTestServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
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
