package persistence

import (
	"fmt"

	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/repository"
	"gorm.io/gorm"
)

type restaurantPersistence struct {
	Conn *gorm.DB
}

var layout = "2006-01-02 15:04:05"

func NewRestaurantPersistence(conn *gorm.DB) repository.RestaurantRepository {
	return &restaurantPersistence{Conn: conn}
}

func (rp *restaurantPersistence) FindAll() (restaurants []*model.Restaurant, err error) {
	db := rp.Conn

	if err := db.Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (rp *restaurantPersistence) Find(id string) (restaurants *model.Restaurant, err error) {
	db := rp.Conn

	if err := db.Where("id = ?", id).First(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (rp *restaurantPersistence) Save(restaurant *model.Restaurant) (*model.Restaurant, error) {

	db := rp.Conn

	if err := db.Create(&restaurant).Error; err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (rp *restaurantPersistence) Update(restaurant *model.Restaurant) (*model.Restaurant, error) {

	db := rp.Conn

	if err := db.Model(&restaurant).Updates(map[string]interface{}{
		"name":          restaurant.Name,
		"fee":           restaurant.Fee,
		"time_required": restaurant.TimeRequired,
	}).Error; err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (rp *restaurantPersistence) Delete(id string) (restaurant *model.Restaurant, err error) {

	db := rp.Conn
	fmt.Println("ここまできてる")

	if err := db.Delete(&model.Restaurant{}, id).Error; err != nil {
		return nil, err
	}

	return restaurant, nil
}
