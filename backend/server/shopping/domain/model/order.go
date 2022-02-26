package model

import "time"

type Order struct {
	ID         string `gorm:"type:id;primary_key;default:uuid_generate_v4()"`
	TotalPrice int32  `gorm:"not null;default 0"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewOrder(order *Order, id string) *Order {
	return &Order{
		ID:         id,
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func (f *Order) GetTotalPrice() int32 {
	return f.TotalPrice
}

func (f *Order) GetCreatedAt() time.Time {
	return f.CreatedAt
}

func (f *Order) GetUpdatedAt() time.Time {
	return f.UpdatedAt
}
