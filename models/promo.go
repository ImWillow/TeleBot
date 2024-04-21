package models

const (
	Promo_URL = "https://www.afk.global/afk-journey-codes"
)

type Promo struct {
	Key    string `json:"key"`
	Reward string `json:"reward"`
}
