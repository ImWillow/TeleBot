package models

const (
	Promo_URL = "https://app-time.ru/games/afk-2-journey/codes"
)

type Promo struct {
	Key    string `json:"key"`
	Reward string `json:"reward"`
	Active bool   `json:"active"`
	Date   string `json:"date"`
}
