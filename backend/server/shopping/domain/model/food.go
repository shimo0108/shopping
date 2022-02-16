package model

import "time"

type Food struct {
	ID           string `gorm:"type:id;primary_key;default:uuid_generate_v4()"`
	RestaurantID string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Price        int32  `gorm:"not null;default 0"`
	Description  string `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewFood(food *Food, id string) *Food {
	return &Food{
		ID:           id,
		RestaurantID: food.RestaurantID,
		Name:         food.Name,
		Price:        food.Price,
		Description:  food.Description,
		CreatedAt:    food.CreatedAt,
		UpdatedAt:    food.UpdatedAt,
	}
}

func (f *Food) GetID() string {
	return f.ID
}

func (f *Food) GetRestaurantID() string {
	return f.RestaurantID
}

func (f *Food) GetName() string {
	return f.Name
}

func (f *Food) GetPrice() int32 {
	return f.Price
}

func (f *Food) GetDescription() string {
	return f.Description
}

func (f *Food) GetCreatedAt() time.Time {
	return f.CreatedAt
}

func (f *Food) GetUpdatedAt() time.Time {
	return f.UpdatedAt
}
