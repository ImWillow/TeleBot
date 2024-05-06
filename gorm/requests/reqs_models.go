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
	GetUserByTelegramID(tID string) (dbmodels.User, error)
	// User Promos
	AddPromosToUser(promos []int64, userID uint) error
	GetUserPromos(userID uint) ([]int64, error)
	// Promos
	GetPromos() ([]dbmodels.Promo, error)
	NewPromo(promo dbmodels.Promo) error
	// ClearPromos() error // NOTE: deprecated
}

type requestModels struct {
	db *gorm.DB
}

func NewRequestModel(db *gorm.DB) RequestModels {
	rm := new(requestModels)
	rm.db = db

	return rm
}
