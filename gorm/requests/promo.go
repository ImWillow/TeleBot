package requests

import (
	dbmodels "telegrambot/gorm/models"
)

func (rm *requestModels) GetPromos() ([]dbmodels.Promo, error) {
	promos := []dbmodels.Promo{}
	if err := rm.db.Find(&promos).Error; err != nil {
		return nil, err
	}

	return promos, nil
}

func (rm *requestModels) NewPromo(promo dbmodels.Promo) error {
	return rm.db.Create(&promo).Error
}

func (rm *requestModels) UpdatePromo(promo dbmodels.Promo) error {
	return rm.db.Model(&promo).Update("sended", promo.Sended).Error
}
