package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Cpf      int `gorm:"unique;not null"`
	Name     string
	Category string
}
