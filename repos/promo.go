package repos

import (
	"telegrambot/gorm"
	"telegrambot/models"
)

type Promo interface {
	GetPromos(userID string) ([]models.Promo, error)
	AddPromos(promos []int64, userNick string) error
}

type promo struct {
	gm gorm.GormModule
}

func NewPromoRepo(gm gorm.GormModule) Promo {
	u := new(promo)
	u.gm = gm

	return u
}

func (p *promo) GetPromos(userID string) ([]models.Promo, error) {
	rm := p.gm.GetRM()
	dbpromos, err := rm.GetPromos()
	if err != nil {
		return nil, err
	}

	user, err := rm.GetUserByTelegramID(userID)
	if err != nil {
		return nil, err
	}

	userPromoIDs, err := rm.GetUserPromos(user.ID)
	if err != nil {
		return nil, err
	}

	var promos []models.Promo
	for _, dbpromo := range dbpromos {
		f := false
		for _, id := range userPromoIDs {
			if uint(id) == dbpromo.ID {
				f = true
				break
			}
		}
		if f {
			continue
		}
		promos = append(promos, models.Promo{
			Key:    dbpromo.Key,
			Reward: dbpromo.Reward,
			Date:   dbpromo.Date,
			Active: dbpromo.Active,
			ID:     dbpromo.ID,
		})
	}

	return promos, nil
}

func (p *promo) AddPromos(promos []int64, userNick string) error {
	rm := p.gm.GetRM()
	u, err := rm.GetUserByTelegramID(userNick)
	if err != nil {
		return err
	}
	if err := rm.AddPromosToUser(promos, u.ID); err != nil {
		return err
	}

	return nil
}
