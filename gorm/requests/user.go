package requests

import (
	dbmodels "telegrambot/gorm/models"
	"telegrambot/models"
)

func (rm *requestModels) NewUser(user models.User) error {
	dbuser := dbmodels.User{
		TelegramID: user.TelegramID,
		Nickname:   user.NickName,
		Role:       user.Role,
	}

	return rm.db.Create(&dbuser).Error
}

func (rm *requestModels) GetUsers() ([]dbmodels.User, error) {
	users := []dbmodels.User{}
	if err := rm.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (rm *requestModels) GetUserByTelegramID(tID string) (dbmodels.User, error) {
	dbUser := dbmodels.User{}
	if err := rm.db.Model(dbmodels.User{TelegramID: tID}).First(&dbUser).Error; err != nil {
		return dbmodels.User{}, err
	}

	return dbUser, nil
}
