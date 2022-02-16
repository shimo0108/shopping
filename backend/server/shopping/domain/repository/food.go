package repository

import (
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
)

type FoodRepository interface {
	FindAll(restaurantID string) ([]*model.Food, error)
	Find(id string) (*model.Food, error)
	Save(*model.Food) (*model.Food, error)
	Update(*model.Food) (*model.Food, error)
	Delete(id string) (*model.Food, error)
}
