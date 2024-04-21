package repos

import "telegrambot/gorm"

type Repos struct {
	UserRepo  User
	PromoRepo Promo
}

func NewRepo(gm gorm.GormModule) Repos {
	var repos Repos
	repos.UserRepo = NewUserRepo(gm)
	repos.PromoRepo = NewPromoRepo(gm)

	return repos
}
