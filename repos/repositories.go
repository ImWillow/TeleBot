package repos

import "telegrambot/gorm"

type Repos struct {
	UserRepo User
}

func NewRepo(gm gorm.GormModule) Repos {
	var repos Repos
	repos.UserRepo = NewUserRepo(gm)

	return repos
}
