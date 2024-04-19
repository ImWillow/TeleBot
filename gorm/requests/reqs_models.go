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
}

type requestModels struct {
	db *gorm.DB
}

func NewRequestModel(db *gorm.DB) RequestModels {
	rm := new(requestModels)
	rm.db = db

	return rm
}
