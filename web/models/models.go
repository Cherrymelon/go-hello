package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model `json:"gorm_._model"`
	Phone      string `json:"phone,omitempty"`
	User       int    `json:"user"`
	Price      int    `json:"price"`
	IsPaid     bool   `json:"is_paid"`
}
