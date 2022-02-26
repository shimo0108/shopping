package handler

import (
	"context"
	"errors"

	pb_order "github.com/shimo0108/shopping/backend/server/pb/order"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/usecase"
)

type OrderHandler struct {
	uc usecase.OrderUsecase
}

func NewOrderHandler(uc usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		uc: uc,
	}
}

func (f *OrderHandler) CreateOrder(ctx context.Context, in *pb_order.CreateOrderRequest) (*pb_order.CreateOrderResponse, error) {

	order, err := f.uc.CreateOrder(&model.Order{}, in.LineFoodIds)
	if err != nil {
		return nil, errors.New("CreateOrders err")
	}
	res := &pb_order.CreateOrderResponse{
		Order: toOrder(order),
	}
	return res, nil
}

func toOrder(order *model.Order) *pb_order.Order {

	res := &pb_order.Order{
		Id:         order.ID,
		TotalPrice: order.TotalPrice,
		CreatedAt:  timestamppb.New(order.CreatedAt),
		UpdatedAt:  timestamppb.New(order.UpdatedAt),
	}
	return res
}
