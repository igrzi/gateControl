package models

import "gorm.io/gorm"

type Access struct {
	gorm.Model
	Cpf  int
	Type string
}

// To know which time the user entered or left the parking lot, you reference the CreatedAt from gorm, and the Tipo to know if it's an entry or a leave.
