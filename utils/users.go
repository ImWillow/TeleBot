package utils

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetUsersFromFile(filepath string) (map[string]string, error) {
	userMap := map[string]string{}
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	users := strings.Split(string(b), ";")
	users = users[:len(users)-1]

	logrus.Debug(users, len(users))
	for _, user := range users {
		userS := strings.Split(user, ":")
		userMap[userS[0]] = userS[1]
	}

	logrus.Debug(userMap)

	return userMap, nil
}

func CheckUserInData(nickname, filepath string) (bool, error) {
	usermap, err := GetUsersFromFile(filepath)
	if err != nil {
		return false, err
	}
	for _, nick := range usermap {
		if nick == nickname {
			return true, nil
		}
	}

	return false, nil
}
