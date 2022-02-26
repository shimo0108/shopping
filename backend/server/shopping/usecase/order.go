package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/repository"
)

type OrderUsecase interface {
	CreateOrder(*model.Order, []string) (*model.Order, error)
}

type orderUsecase struct {
	repo repository.OrderRepository
}

func NewOrderUsecase(repo repository.OrderRepository) OrderUsecase {
	return &orderUsecase{
		repo: repo,
	}
}

func (r *orderUsecase) CreateOrder(Order *model.Order, ids []string) (*model.Order, error) {
	uuidObj, _ := uuid.NewUUID()

	newOrder := model.NewOrder(Order, uuidObj.String())
	res, err := r.repo.Save(newOrder, ids)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
