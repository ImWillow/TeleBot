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

func (rm *requestModels) AddPromoToUser(promo string, userID uint) error {
	if err := rm.db.Model(&dbmodels.UserPromos{}).Where("user_id = ?", userID).Update("promos", []string{promo}).Error; err != nil {
		return err
	}
	return nil
}

func (rm *requestModels) GetUserPromos(userID uint) ([]string, error) {
	promos := []string{}
	if err := rm.db.Raw(`select promos from user_promos where user_id = ?`, userID).Find(&promos).Error; err != nil {
		return nil, err
	}
	return promos, nil
}
