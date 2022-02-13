package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/repository"
)

type RestaurantUsecase interface {
	ListRestaurants() ([]*model.Restaurant, error)
	GetRestaurant(id string) (*model.Restaurant, error)
	CreateRestaurant(*model.Restaurant) (*model.Restaurant, error)
	UpdateRestaurant(*model.Restaurant) (*model.Restaurant, error)
	DeleteRestaurant(id string) (*model.Restaurant, error)
}

type restaurantUsecase struct {
	repo repository.RestaurantRepository
}

func NewRestaurantUsecase(repo repository.RestaurantRepository) RestaurantUsecase {
	return &restaurantUsecase{
		repo: repo,
	}
}

func (r *restaurantUsecase) ListRestaurants() ([]*model.Restaurant, error) {
	restaurants, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *restaurantUsecase) GetRestaurant(id string) (*model.Restaurant, error) {

	res, err := r.repo.Find(id)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *restaurantUsecase) CreateRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error) {
	uuidObj, _ := uuid.NewUUID()

	newRestaurant := model.NewRestaurant(restaurant, uuidObj.String())
	res, err := r.repo.Save(newRestaurant)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *restaurantUsecase) UpdateRestaurant(restaurant *model.Restaurant) (*model.Restaurant, error) {

	res, err := r.repo.Update(restaurant)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *restaurantUsecase) DeleteRestaurant(id string) (*model.Restaurant, error) {

	res, err := r.repo.Delete(id)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
