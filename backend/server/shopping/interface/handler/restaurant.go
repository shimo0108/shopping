package handler

import (
	"context"
	"errors"

	pb_restaurant "github.com/shimo0108/shopping/backend/server/pb/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/usecase"
)

type RestaurantHandler struct {
	uc usecase.RestaurantUsecase
}

func NewRestaurantHandler(uc usecase.RestaurantUsecase) *RestaurantHandler {
	return &RestaurantHandler{
		uc: uc,
	}
}

func (h *RestaurantHandler) ListRestaurants(ctx context.Context, in *pb_restaurant.ListRestaurantRequest) (*pb_restaurant.ListRestaurantResponse, error) {
	restaurants, err := h.uc.ListRestaurants()
	if err != nil {
		return nil, errors.New("ListRestaurants err")
	}
	res := &pb_restaurant.ListRestaurantResponse{
		Restaurants: toRestaurants(restaurants),
	}
	return res, nil
}

func (h *RestaurantHandler) GetRestaurant(ctx context.Context, in *pb_restaurant.ReadRestaurantRequest) (*pb_restaurant.ReadRestaurantResponse, error) {
	restaurant, err := h.uc.GetRestaurant(in.GetId())
	if err != nil {
		return nil, errors.New("CreateRestaurants err")
	}
	res := &pb_restaurant.ReadRestaurantResponse{
		Restaurant: toRestaurant(restaurant),
	}
	return res, nil
}

func (h *RestaurantHandler) CreateRestaurant(ctx context.Context, in *pb_restaurant.CreateRestaurantRequest) (*pb_restaurant.CreateRestaurantResponse, error) {
	grpcReq := transformNewRestaurantData(in)

	restaurant, err := h.uc.CreateRestaurant(grpcReq)
	if err != nil {
		return nil, errors.New("CreateRestaurants err")
	}
	res := &pb_restaurant.CreateRestaurantResponse{
		Restaurant: toRestaurant(restaurant),
	}
	return res, nil
}

func (h *RestaurantHandler) UpdateRestaurant(ctx context.Context, in *pb_restaurant.UpdateRestaurantRequest) (*pb_restaurant.UpdateRestaurantResponse, error) {
	grpcReq := transformRestaurantData(in)

	restaurant, err := h.uc.UpdateRestaurant(grpcReq)
	if err != nil {
		return nil, errors.New("UpdateRestaurants err")
	}
	res := &pb_restaurant.UpdateRestaurantResponse{
		Restaurant: toRestaurant(restaurant),
	}
	return res, nil
}

func (h *RestaurantHandler) DeleteRestaurant(ctx context.Context, in *pb_restaurant.DeleteRestaurantRequest) (*pb_restaurant.DeleteRestaurantResponse, error) {

	restaurant, err := h.uc.DeleteRestaurant(in.GetId())
	if err != nil {
		return nil, errors.New("DeleteRestaurants err")
	}
	res := &pb_restaurant.DeleteRestaurantResponse{
		Restaurant: toRestaurant(restaurant),
	}
	return res, nil
}

func toRestaurants(restaurants []*model.Restaurant) []*pb_restaurant.Restaurant {
	res := make([]*pb_restaurant.Restaurant, len(restaurants))
	for i, restaurant := range restaurants {
		res[i] = &pb_restaurant.Restaurant{
			Id:           restaurant.ID,
			Name:         restaurant.Name,
			Fee:          restaurant.Fee,
			TimeRequired: restaurant.TimeRequired,
			CreatedAt:    timestamppb.New(restaurant.CreatedAt),
			UpdatedAt:    timestamppb.New(restaurant.UpdatedAt),
		}
	}
	return res
}

func toRestaurant(restaurant *model.Restaurant) *pb_restaurant.Restaurant {

	res := &pb_restaurant.Restaurant{
		Id:           restaurant.ID,
		Name:         restaurant.Name,
		Fee:          restaurant.Fee,
		TimeRequired: restaurant.TimeRequired,
		CreatedAt:    timestamppb.New(restaurant.CreatedAt),
		UpdatedAt:    timestamppb.New(restaurant.UpdatedAt),
	}
	return res
}

func transformNewRestaurantData(restaurant *pb_restaurant.CreateRestaurantRequest) *model.Restaurant {
	res := &model.Restaurant{
		Name:         restaurant.Name,
		Fee:          restaurant.Fee,
		TimeRequired: restaurant.TimeRequired,
	}
	return res
}

func transformRestaurantData(restaurant *pb_restaurant.UpdateRestaurantRequest) *model.Restaurant {

	res := &model.Restaurant{
		ID:           restaurant.Id,
		Name:         restaurant.Name,
		Fee:          restaurant.Fee,
		TimeRequired: restaurant.TimeRequired,
	}
	return res
}
