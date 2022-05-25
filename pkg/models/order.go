package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	// Id        uint64 `gorm:"primaryKey"`
	Price     int64 `json:"price"`
	ProductId int64 `json:"product_id"`
	UserId    int64 `json:"user_id"`
}
