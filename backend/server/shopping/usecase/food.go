package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/repository"
)

type FoodUsecase interface {
	ListFoods(restaurantId string) ([]*model.Food, error)
	GetFood(id string) (*model.Food, error)
	CreateFood(*model.Food) (*model.Food, error)
	UpdateFood(*model.Food) (*model.Food, error)
	DeleteFood(id string) (*model.Food, error)
}

type foodUsecase struct {
	repo repository.FoodRepository
}

func NewFoodUsecase(repo repository.FoodRepository) FoodUsecase {
	return &foodUsecase{
		repo: repo,
	}
}

func (r *foodUsecase) ListFoods(restaurantId string) ([]*model.Food, error) {
	Foods, err := r.repo.FindAll(restaurantId)
	if err != nil {
		return nil, err
	}
	return Foods, nil
}

func (r *foodUsecase) GetFood(id string) (*model.Food, error) {

	res, err := r.repo.Find(id)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *foodUsecase) CreateFood(food *model.Food) (*model.Food, error) {
	uuidObj, _ := uuid.NewUUID()

	newFood := model.NewFood(food, uuidObj.String())
	res, err := r.repo.Save(newFood)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *foodUsecase) UpdateFood(food *model.Food) (*model.Food, error) {

	res, err := r.repo.Update(food)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *foodUsecase) DeleteFood(id string) (*model.Food, error) {

	res, err := r.repo.Delete(id)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
