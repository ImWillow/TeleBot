package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TelegramID string
	Nickname   string
	Role       string
}

type UserPromos struct {
	gorm.Model
	UserID   uint
	PromoIDs pq.Int64Array `gorm:"type:integer[]"`
}
