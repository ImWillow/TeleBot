package models

import "gorm.io/gorm"

type Promo struct {
	gorm.Model
	Key    string
	Reward string
	Sended bool
}
