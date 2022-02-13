package repository

import (
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
)

type RestaurantRepository interface {
	FindAll() ([]*model.Restaurant, error)
	Find(id string) (*model.Restaurant, error)
	Save(*model.Restaurant) (*model.Restaurant, error)
	Update(*model.Restaurant) (*model.Restaurant, error)
	Delete(id string) (*model.Restaurant, error)
}
