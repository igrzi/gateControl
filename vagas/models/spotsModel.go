package models

import "gorm.io/gorm"

type Spots struct {
	gorm.Model
	QuantityAvailable int
	MaxQuantity       int
}
