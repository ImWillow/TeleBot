package utils

import (
	dbmodels "telegrambot/gorm/models"
	"telegrambot/models"
)

func UsersFromDB(dbuser []dbmodels.User) []models.User {
	users := []models.User{}
	for _, dbu := range dbuser {
		users = append(users, models.User{
			TelegramID: dbu.TelegramID,
			NickName:   dbu.Nickname,
			Role:       dbu.Role,
		})
	}

	return users
}
