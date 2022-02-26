package repository

import (
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
)

type OrderRepository interface {
	Save(*model.Order, []string) (*model.Order, error)
}
