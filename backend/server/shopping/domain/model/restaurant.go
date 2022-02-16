package model

import (
	"time"
)

type Restaurant struct {
	ID           string `gorm:"type:id;primary_key;default:uuid_generate_v4()"`
	Name         string `gorm:"not null"`
	Fee          int32  `gorm:"not null;default 0"`
	TimeRequired int32  `gorm:"not null;default 0"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewRestaurant(restaurant *Restaurant, uid string) *Restaurant {
	return &Restaurant{
		ID:           uid,
		Name:         restaurant.Name,
		Fee:          restaurant.Fee,
		TimeRequired: restaurant.TimeRequired,
		CreatedAt:    restaurant.CreatedAt,
		UpdatedAt:    restaurant.UpdatedAt,
	}
}

func (r *Restaurant) GetID() string {
	return r.ID
}

func (r *Restaurant) GetName() string {
	return r.Name
}

func (r *Restaurant) GetFee() int32 {
	return r.Fee
}

func (r *Restaurant) GetTimeRequired() int32 {
	return r.TimeRequired
}

func (r *Restaurant) GetCreatedAt() time.Time {
	return r.CreatedAt
}

func (r *Restaurant) GetUpdatedAt() time.Time {
	return r.UpdatedAt
}
