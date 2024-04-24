package requests

import (
	dbmodels "telegrambot/gorm/models"
	"telegrambot/models"

	"gorm.io/gorm"
)

type RequestModels interface {
	// User
	NewUser(user models.User) error
	GetUsers() ([]dbmodels.User, error)
	// User Promos
	AddPromoToUser(promo string, userID uint) error
	GetUserPromos(userID uint) ([]string, error)
	// Promos
	GetPromos() ([]dbmodels.Promo, error)
	NewPromo(promo dbmodels.Promo) error
	ClearPromos() error
}

type requestModels struct {
	db *gorm.DB
}

func NewRequestModel(db *gorm.DB) RequestModels {
	rm := new(requestModels)
	rm.db = db

	return rm
}
