package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/repository"
)

type LineFoodUsecase interface {
	ListLineFoods() (*model.LineFoodResult, error)
	CreateLineFood(*model.LineFood) (*model.LineFood, error)
	ReplaceLineFood(restaurantID string) ([]*model.LineFood, error)
}

type lineFoodUsecase struct {
	repo repository.LineFoodRepository
}

func NewLineFoodUsecase(repo repository.LineFoodRepository) LineFoodUsecase {
	return &lineFoodUsecase{
		repo: repo,
	}
}

func (lf *lineFoodUsecase) ListLineFoods() (*model.LineFoodResult, error) {
	lineFoods, err := lf.repo.FindActiveLineFoods()

	if err != nil {
		return nil, err
	}
	return lineFoods, nil
}

func (lf *lineFoodUsecase) CreateLineFood(lineFood *model.LineFood) (*model.LineFood, error) {
	uuidObj, _ := uuid.NewUUID()

	newLineFood := model.NewLineFood(lineFood, uuidObj.String())
	res, err := lf.repo.Save(newLineFood)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (lf *lineFoodUsecase) ReplaceLineFood(restaurantId string) ([]*model.LineFood, error) {

	res, err := lf.repo.Replace(restaurantId)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
