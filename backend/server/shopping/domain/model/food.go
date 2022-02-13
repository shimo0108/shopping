package model

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	RestaurantID int32  `gorm:"not null"`
	Name         string `gorm:"not null"`
	Price        int32  `gorm:"not null;default 0"`
	Description  string `gorm:"type:varchar(255);not null"`
}
