package migs

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M2 *gormigrate.Migration = &gormigrate.Migration{
	ID: "2",
	Migrate: func(tx *gorm.DB) error {
		type promo struct {
			gorm.Model
			Key    string
			Reward string
			Sended bool
		}
		return tx.Migrator().CreateTable(&promo{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("promos")
	},
}
