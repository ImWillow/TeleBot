package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"telegrambot/models"
)

func GetUsersFromFile() ([]models.User, error) {
	users := []models.User{}
	b, err := os.ReadFile(models.Path_Members)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func CheckUserInData(user models.User) (bool, error) {
	users, err := GetUsersFromFile()
	if err != nil {
		return false, err
	}
	for _, u := range users {
		if user.NickName == u.NickName || user.TelegramID == u.TelegramID {
			return true, nil
		}
	}

	return false, nil
}

func AddUserToData(user models.User) error {
	already, err := CheckUserInData(user)
	if err != nil {
		return fmt.Errorf("can't check user, error: '%s'", err.Error())
	}
	if !already {
		users, err := GetUsersFromFile()
		if err != nil {
			return err
		}
		users = append(users, user)
		b, err := json.Marshal(users)
		if err != nil {
			return err
		}
		if err = os.WriteFile(models.Path_Members, b, 0644); err != nil {
			return err
		}

		return nil
	}

	return errors.New("user already in data file")
}

func GetUserByNickname(nickname string) (models.User, error) {
	users, err := GetUsersFromFile()
	if err != nil {
		return models.User{}, err
	}
	for _, u := range users {
		if u.NickName == nickname {
			return u, nil
		}
	}

	return models.User{}, errors.New("user not found")
}

func AddUserRole(user models.User, role string) error {
	roleModels := []models.Role{}
	b, err := os.ReadFile(models.Path_ID)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &roleModels); err != nil {
		return err
	}

	for _, r := range roleModels {
		if r.TelegramID == user.TelegramID {
			r.RoleType = role

			return nil
		}
	}

	roleModels = append(roleModels, models.Role{
		TelegramID: user.TelegramID,
		RoleType:   role,
	})

	b, err = json.Marshal(roleModels)
	if err != nil {
		return err
	}
	if err = os.WriteFile(models.Path_ID, b, 0644); err != nil {
		return err
	}

	return nil
}

func GetUserRole(user models.User) (string, error) {
	role := []models.Role{}
	b, err := os.ReadFile(models.Path_ID)
	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(b, &role); err != nil {
		return "", err
	}

	for _, r := range role {
		if r.TelegramID == user.TelegramID {
			return r.RoleType, nil
		}
	}

	return "", errors.New("user not found")
}
