package handler

import (
	"context"
	"errors"

	pb_food "github.com/shimo0108/shopping/backend/server/pb/food"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/usecase"
)

type FoodHandler struct {
	uc usecase.FoodUsecase
}

func NewFoodHandler(uc usecase.FoodUsecase) *FoodHandler {
	return &FoodHandler{
		uc: uc,
	}
}

func (f *FoodHandler) ListFoods(ctx context.Context, in *pb_food.ListFoodRequest) (*pb_food.ListFoodResponse, error) {
	Foods, err := f.uc.ListFoods(in.GetRestaurantId())

	if err != nil {
		return nil, errors.New("ListFoods err")
	}
	res := &pb_food.ListFoodResponse{
		Foods: toFoods(Foods),
	}
	return res, nil
}

func (f *FoodHandler) GetFood(ctx context.Context, in *pb_food.ReadFoodRequest) (*pb_food.ReadFoodResponse, error) {
	Food, err := f.uc.GetFood(in.GetId())
	if err != nil {
		return nil, errors.New("CreateFoods err")
	}
	res := &pb_food.ReadFoodResponse{
		Food: toFood(Food),
	}
	return res, nil
}

func (f *FoodHandler) CreateFood(ctx context.Context, in *pb_food.CreateFoodRequest) (*pb_food.CreateFoodResponse, error) {
	grpcReq := transformNewFoodData(in)

	Food, err := f.uc.CreateFood(grpcReq)
	if err != nil {
		return nil, errors.New("CreateFoods err")
	}
	res := &pb_food.CreateFoodResponse{
		Food: toFood(Food),
	}
	return res, nil
}

func (f *FoodHandler) UpdateFood(ctx context.Context, in *pb_food.UpdateFoodRequest) (*pb_food.UpdateFoodResponse, error) {
	grpcReq := transformFoodData(in)

	Food, err := f.uc.UpdateFood(grpcReq)
	if err != nil {
		return nil, errors.New("UpdateFoods err")
	}
	res := &pb_food.UpdateFoodResponse{
		Food: toFood(Food),
	}
	return res, nil
}

func (f *FoodHandler) DeleteFood(ctx context.Context, in *pb_food.DeleteFoodRequest) (*pb_food.DeleteFoodResponse, error) {

	Food, err := f.uc.DeleteFood(in.GetId())
	if err != nil {
		return nil, errors.New("DeleteFoods err")
	}
	res := &pb_food.DeleteFoodResponse{
		Food: toFood(Food),
	}
	return res, nil
}

func toFoods(Foods []*model.Food) []*pb_food.Food {
	res := make([]*pb_food.Food, len(Foods))
	for i, Food := range Foods {
		res[i] = &pb_food.Food{
			Id:           Food.ID,
			RestaurantId: Food.RestaurantID,
			Name:         Food.Name,
			Price:        Food.Price,
			Description:  Food.Description,
			CreatedAt:    timestamppb.New(Food.CreatedAt),
			UpdatedAt:    timestamppb.New(Food.UpdatedAt),
		}
	}
	return res
}

func toFood(Food *model.Food) *pb_food.Food {

	res := &pb_food.Food{
		Id:          Food.ID,
		Name:        Food.Name,
		Price:       Food.Price,
		Description: Food.Description,
		CreatedAt:   timestamppb.New(Food.CreatedAt),
		UpdatedAt:   timestamppb.New(Food.UpdatedAt),
	}
	return res
}

func transformNewFoodData(Food *pb_food.CreateFoodRequest) *model.Food {
	res := &model.Food{
		RestaurantID: Food.RestaurantId,
		Name:         Food.Name,
		Price:        Food.Price,
		Description:  Food.Description,
	}
	return res
}

func transformFoodData(Food *pb_food.UpdateFoodRequest) *model.Food {

	res := &model.Food{
		ID:          Food.Id,
		Name:        Food.Name,
		Price:       Food.Price,
		Description: Food.Description,
	}
	return res
}
