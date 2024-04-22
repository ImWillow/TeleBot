package repos

import (
	"errors"
	"telegrambot/gorm"
	"telegrambot/models"
	"telegrambot/utils"
)

type User interface {
	NewUser(user models.User) error
	GetUsers() ([]models.User, error)
}

type user struct {
	gm gorm.GormModule
}

func NewUserRepo(gm gorm.GormModule) User {
	u := new(user)
	u.gm = gm

	return u
}

func (u *user) NewUser(user models.User) error {
	rm := u.gm.GetRM()

	dbusers, err := rm.GetUsers()
	if err != nil {
		return err
	}
	users := utils.UsersFromDB(dbusers)
	if !u.CheckUserInUsers(user, users) {
		if err := rm.NewUser(user); err != nil {
			return err
		}

		return nil
	}

	return errors.New("user already in DB")
}

func (u *user) GetUsers() ([]models.User, error) {
	rm := u.gm.GetRM()

	dbusers, err := rm.GetUsers()
	if err != nil {
		return nil, err
	}
	users := utils.UsersFromDB(dbusers)

	return users, nil
}

func (u *user) CheckUserInUsers(user models.User, users []models.User) bool {
	for _, u := range users {
		if user.NickName == u.NickName || user.TelegramID == u.TelegramID {
			return true
		}
	}

	return false
}
