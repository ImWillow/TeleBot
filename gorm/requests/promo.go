package requests

import (
	dbmodels "telegrambot/gorm/models"
	"telegrambot/models"
)

func (rm *requestModels) GetPromos() ([]dbmodels.Promo, error) {
	promos := []dbmodels.Promo{}
	if err := rm.db.Find(&promos).Error; err != nil {
		return nil, err
	}

	return promos, nil
}

func (rm *requestModels) NewPromo(promo models.Promo) error {
	dbpromo := dbmodels.Promo{
		Key:    promo.Key,
		Reward: promo.Reward,
	}

	return rm.db.Create(&dbpromo).Error
}
