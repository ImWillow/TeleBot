package models

// Template
const (
	UserTemplate = "@%s:%s;"
)

const (
	Role_admin  = "admin"
	Role_member = "member"
)

type User struct {
	TelegramID string `json:"telegramId"`
	NickName   string `json:"nickname"`
}

type Role struct {
	TelegramID string `json:"telegramId"`
	RoleType   string `json:"roleType"`
}
