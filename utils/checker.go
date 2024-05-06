package utils

import (
	dbmodels "telegrambot/gorm/models"
)

func CheckPromo(oldPromos, newPromos []dbmodels.Promo) []dbmodels.Promo {
	var truePromos []dbmodels.Promo
	for _, newPromo := range newPromos {
		f := false
		for _, oldPromo := range oldPromos {
			if newPromo.Key == oldPromo.Key {
				f = true
				break
			}
		}
		if f {
			continue
		}
		truePromos = append(truePromos, newPromo)
	}

	return truePromos
}
