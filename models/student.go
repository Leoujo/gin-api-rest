package models

import (
	"gorm.io/gorm"
)

// Student model
type Student struct {
	gorm.Model
	Name string `json:"nome"`
	CPF string `json:"cpf"`
	RG string `json:"rg"`
}

// Every time we add gorm.Model to the struct it automatically creates:
// ID, CreatedAt, UpdatedAt, DeletedAt, Name