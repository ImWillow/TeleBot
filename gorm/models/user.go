package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TelegramID string
	Nickname   string
	Role       string
}
