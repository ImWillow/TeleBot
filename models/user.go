package models

const (
	Role_admin  = "admin"
	Role_member = "member"
)

type User struct {
	TelegramID string `json:"telegramID"`
	NickName   string `json:"nickname"`
	Role       string `json:"role"`
}
