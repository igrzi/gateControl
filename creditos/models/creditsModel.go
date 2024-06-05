package models

import "gorm.io/gorm"

type Credits struct {
	gorm.Model
	Cpf          int `gorm:"unique;not null"`
	CreditAmount int
}
