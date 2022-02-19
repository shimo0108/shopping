package handler

import (
	"context"
	"errors"

	pb_line_food "github.com/shimo0108/shopping/backend/server/pb/line_food"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/usecase"
)

type LineFoodHandler struct {
	uc usecase.LineFoodUsecase
}

func NewLineFoodHandler(uc usecase.LineFoodUsecase) *LineFoodHandler {
	return &LineFoodHandler{
		uc: uc,
	}
}

func (h *LineFoodHandler) ListLineFoods(ctx context.Context, in *pb_line_food.ListLineFoodRequest) (*pb_line_food.ListLineFoodResponse, error) {
	lineFoods, err := h.uc.ListLineFoods()
	if err != nil {
		return nil, errors.New("ListLineFoods err")
	}
	res := &pb_line_food.ListLineFoodResponse{
		LineFoodIds: lineFoods.LineFoodIds,
		Restaurant:  toRestaurant(lineFoods.Restaurant),
		Count:       lineFoods.Count,
		Amount:      lineFoods.Amount,
	}
	return res, nil
}

func (h *LineFoodHandler) CreateLineFood(ctx context.Context, in *pb_line_food.CreateLineFoodRequest) (*pb_line_food.CreateLineFoodResponse, error) {
	grpcReq := transformNewLineFoodData(in)

	lineFood, err := h.uc.CreateLineFood(grpcReq)
	if err != nil {
		return nil, errors.New("CreateLineFoods err")
	}
	res := &pb_line_food.CreateLineFoodResponse{
		LineFood: toLineFood(lineFood),
	}
	return res, nil
}

func (h *LineFoodHandler) ReplaceLineFood(ctx context.Context, in *pb_line_food.ReplaceLineFoodRequest) (*pb_line_food.ReplaceLineFoodResponse, error) {

	lineFoods, err := h.uc.ReplaceLineFood(in.GetRestaurantId())
	if err != nil {
		return nil, errors.New("ReplaceLineFoods err")
	}
	res := &pb_line_food.ReplaceLineFoodResponse{
		LineFoods: toLineFoods(lineFoods),
	}
	return res, nil
}

func toLineFoods(lineFoods []*model.LineFood) []*pb_line_food.LineFood {
	res := make([]*pb_line_food.LineFood, len(lineFoods))
	for i, lineFood := range lineFoods {
		res[i] = &pb_line_food.LineFood{
			Id:           lineFood.ID,
			RestaurantId: lineFood.RestaurantID,
			FoodId:       lineFood.FoodID,
			OrderId:      lineFood.OrderID,
			Count:        lineFood.Count,
			Active:       lineFood.Active,
			CreatedAt:    timestamppb.New(lineFood.CreatedAt),
			UpdatedAt:    timestamppb.New(lineFood.UpdatedAt),
		}
	}
	return res
}

func toLineFood(lineFood *model.LineFood) *pb_line_food.LineFood {

	res := &pb_line_food.LineFood{
		Id:           lineFood.ID,
		RestaurantId: lineFood.RestaurantID,
		FoodId:       lineFood.FoodID,
		OrderId:      lineFood.OrderID,
		Count:        lineFood.Count,
		Active:       lineFood.Active,
		CreatedAt:    timestamppb.New(lineFood.CreatedAt),
		UpdatedAt:    timestamppb.New(lineFood.UpdatedAt),
	}
	return res
}

func toLineFoodResult(lineFoodResult *model.LineFoodResult) *pb_line_food.ListLineFoodResponse {

	res := &pb_line_food.ListLineFoodResponse{
		LineFoodIds: lineFoodResult.LineFoodIds,
		Restaurant:  toRestaurant(lineFoodResult.Restaurant),
		Count:       lineFoodResult.Count,
		Amount:      lineFoodResult.Amount,
	}
	return res
}

func transformNewLineFoodData(lineFood *pb_line_food.CreateLineFoodRequest) *model.LineFood {
	res := &model.LineFood{
		RestaurantID: lineFood.RestaurantId,
		FoodID:       lineFood.FoodId,
		Count:        lineFood.Count,
	}
	return res
}
