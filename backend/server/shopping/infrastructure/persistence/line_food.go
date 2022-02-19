package persistence

import (
	"github.com/shimo0108/shopping/backend/server/shopping/domain/model"
	"github.com/shimo0108/shopping/backend/server/shopping/domain/repository"
	"gorm.io/gorm"
)

type lineFoodPersistence struct {
	Conn *gorm.DB
}

func NewLineFoodPersistence(conn *gorm.DB) repository.LineFoodRepository {
	return &lineFoodPersistence{Conn: conn}
}

func (lf *lineFoodPersistence) FindActiveLineFoods() (lineFoodResult *model.LineFoodResult, err error) {
	db := lf.Conn
	var ids []string
	var restaurant *model.Restaurant
	var sum int32
	var totalAmount int32
	var lineFoods []model.LineFood

	db.Model(&lineFoods).Where("active = ?", true).Pluck("id", &ids)

	db.Where("active = ?", true).Find(&lineFoods)
	for i := range lineFoods {
		db.Model(lineFoods[i]).Association("Food").Find(&lineFoods[i].Food)
	}

	if err := db.Where("id = ?", lineFoods[0].RestaurantID).Find(&restaurant).Error; err != nil {
		return nil, err
	}
	for _, lf := range lineFoods {
		sum += lf.Count
		totalAmount += lf.Food.Price * lf.Count
	}
	lineFoodResult = model.NewLineFoodResult(ids, restaurant, sum, totalAmount)

	return lineFoodResult, nil
}

func (lf *lineFoodPersistence) Save(lineFood *model.LineFood) (*model.LineFood, error) {

	db := lf.Conn
	initLineFood := setLineFood(db, lineFood)

	if err := db.Save(&initLineFood).Error; err != nil {
		return nil, err
	}
	return lineFood, nil
}

func (lf *lineFoodPersistence) Replace(restaurantId string) (lineFoods []*model.LineFood, err error) {

	db := lf.Conn

	if err := db.Model(&lineFoods).Not("restaurant_id = ?", restaurantId).Updates(map[string]interface{}{
		"active": false,
	}).Error; err != nil {
		return nil, err
	}
	if err := db.Where("active = ?", true).Find(&lineFoods).Error; err != nil {
		return nil, err
	}

	return lineFoods, nil
}

func setLineFood(db *gorm.DB, lineFood *model.LineFood) *model.LineFood {
	var foodCount int64
	var existingLine *model.LineFood
	db.Where("food_id = ?", lineFood.FoodID).First(&existingLine).Count(&foodCount)

	if foodCount == 0 {
		lineFood.Active = true
		return lineFood
	} else {
		existingLine.Active = true
		existingLine.Count += lineFood.Count
		return existingLine
	}

}
