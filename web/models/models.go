package models

import (
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
	"time"
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
	Sessionid guuid.UUID `gorm:"primaryKey" json:"Sessionid"`
	Expires   time.Time  `json:"-"`
	UserRefer uint       `json:"-"`
	CreatedAt int64      `gorm:"autoCreateTime" json:"-" `
}
