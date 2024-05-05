package repos

import (
	"telegrambot/gorm"
	"telegrambot/models"
)

type Promo interface {
	GetPromos() ([]models.Promo, error)
}

type promo struct {
	gm gorm.GormModule
}

func NewPromoRepo(gm gorm.GormModule) Promo {
	u := new(promo)
	u.gm = gm

	return u
}

func (p *promo) GetPromos() ([]models.Promo, error) {
	rm := p.gm.GetRM()
	dbpromos, err := rm.GetPromos()
	if err != nil {
		return nil, err
	}

	var promos []models.Promo
	for _, dbpromo := range dbpromos {
		promos = append(promos, models.Promo{
			Key:    dbpromo.Key,
			Reward: dbpromo.Reward,
			Date:   dbpromo.Date,
			Active: dbpromo.Active,
		})
	}

	return promos, nil
}
