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

func (rm *requestModels) AddPromosToUser(promos []int64, userID uint) error {
	var uPromo dbmodels.UserPromos
	if err := rm.db.Where(dbmodels.UserPromos{UserID: userID}).FirstOrCreate(&uPromo).Error; err != nil {
		return err
	}
	for _, p := range promos {
		f := false
		for _, dbp := range uPromo.PromoIDs {
			if dbp == p {
				f = true
				break
			}
		}
		if f {
			continue
		}

		uPromo.PromoIDs = append(uPromo.PromoIDs, p)
	}

	if err := rm.db.Save(&uPromo).Error; err != nil {
		return err
	}

	return nil
}

func (rm *requestModels) GetUserPromos(userID uint) ([]int64, error) {
	promos := dbmodels.UserPromos{}
	if err := rm.db.Model(dbmodels.UserPromos{UserID: userID}).First(&promos).Error; err != nil {
		return nil, err
	}
	promosInt64 := []int64{}
	promosInt64 = append(promosInt64, promos.PromoIDs...)
	return promosInt64, nil
}

func (rm *requestModels) GetUserByTelegramID(tID string) (dbmodels.User, error) {
	dbUser := dbmodels.User{}
	if err := rm.db.Model(dbmodels.User{TelegramID: tID}).First(&dbUser).Error; err != nil {
		return dbmodels.User{}, err
	}

	return dbUser, nil
}
