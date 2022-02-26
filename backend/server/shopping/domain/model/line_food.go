package model

import "time"

type LineFood struct {
	ID           string `gorm:"type:id;primary_key;default:uuid_generate_v4()"`
	RestaurantID string `gorm:"not null"`
	FoodID       string `gorm:"not null"`
	OrderID      string
	Count        int32      `gorm:"not null;default 0"`
	Active       bool       `gorm:"not null;default false"`
	Restaurant   Restaurant `gorm:"ForeignKey:RestaurantID;AssociationForeignKey:ID"`
	Food         Food       `gorm:"ForeignKey:FoodID;AssociationForeignKey:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type LineFoodResult struct {
	LineFoodIds []string
	Restaurant  *Restaurant
	Count       int32
	Amount      int32
}

func NewLineFood(lineFood *LineFood, id string) *LineFood {
	return &LineFood{
		ID:           id,
		RestaurantID: lineFood.RestaurantID,
		FoodID:       lineFood.FoodID,
		OrderID:      lineFood.OrderID,
		Count:        lineFood.Count,
		Active:       lineFood.Active,
		CreatedAt:    lineFood.CreatedAt,
		UpdatedAt:    lineFood.UpdatedAt,
	}
}

func NewLineFoodResult(ids []string, restaurant *Restaurant, sum int32, totalAmount int32) *LineFoodResult {
	return &LineFoodResult{
		LineFoodIds: ids,
		Restaurant:  restaurant,
		Count:       sum,
		Amount:      totalAmount,
	}
}

func (lf *LineFood) GetID() string {
	return lf.ID
}

func (lf *LineFood) GetOrderID() string {
	return lf.OrderID
}

func (lf *LineFood) GetCount() int32 {
	return lf.Count
}

func (lf *LineFood) GetActive() bool {
	return lf.Active
}

func (lf *LineFood) GetCreatedAt() time.Time {
	return lf.CreatedAt
}

func (lf *LineFood) GetUpdatedAt() time.Time {
	return lf.UpdatedAt
}
