package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TelegramID string
	Nickname   string
	Role       string
}

type UserPromos struct {
	gorm.Model
	UserID uint
	Promos []string `gorm:"type:text[]"`
}
