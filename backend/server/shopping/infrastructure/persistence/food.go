package persistence

import (
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/repository"
	"gorm.io/gorm"
)

type foodPersistence struct {
	Conn *gorm.DB
}

func NewFoodPersistence(conn *gorm.DB) repository.FoodRepository {
	return &foodPersistence{Conn: conn}
}

func (rp *foodPersistence) FindAll(restaurantId string) (foods []*model.Food, err error) {
	db := rp.Conn

	if err := db.Where("restaurant_id = ?", restaurantId).Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

func (rp *foodPersistence) Find(id string) (foods *model.Food, err error) {
	db := rp.Conn

	if err := db.Where("id = ?", id).First(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

func (rp *foodPersistence) Save(food *model.Food) (*model.Food, error) {

	db := rp.Conn

	if err := db.Create(&food).Error; err != nil {
		return nil, err
	}
	return food, nil
}

func (rp *foodPersistence) Update(food *model.Food) (*model.Food, error) {

	db := rp.Conn

	if err := db.Model(&food).Updates(map[string]interface{}{
		"name":        food.Name,
		"price":       food.Price,
		"description": food.Description,
	}).Error; err != nil {
		return nil, err
	}
	return food, nil
}

func (rp *foodPersistence) Delete(id string) (food *model.Food, err error) {

	db := rp.Conn

	if err := db.Where("id = ?", id).Delete(&food).Error; err != nil {
		return nil, err
	}

	return food, nil
}
