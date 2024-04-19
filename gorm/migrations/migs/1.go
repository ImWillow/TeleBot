package migs

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M1 *gormigrate.Migration = &gormigrate.Migration{
	ID: "1",
	Migrate: func(tx *gorm.DB) error {
		type user struct {
			gorm.Model
			TelegramID string
			Nickname   string
			Role       string
		}
		return tx.Migrator().CreateTable(&user{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
