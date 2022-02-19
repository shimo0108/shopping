package repository

import (
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
)

type LineFoodRepository interface {
	FindActiveLineFoods() (*model.LineFoodResult, error)
	Save(*model.LineFood) (*model.LineFood, error)
	Replace(restaurantId string) ([]*model.LineFood, error)
}
