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

type User struct {
	gorm.Model `json:"gorm_._model"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	Name       string `json:"name" gorm:"unique"`
}

type Session struct {
	gorm.Model `json:"gorm_._model"`
	SessionKey string `json:"session_key"`
	UserId     int    `json:"user_id"`
	ExpiresAt  int    `json:"expires_at"`
}
