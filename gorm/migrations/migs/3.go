package migs

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M3 *gormigrate.Migration = &gormigrate.Migration{
	ID: "3",
	Migrate: func(tx *gorm.DB) error {
		type userPromos struct {
			gorm.Model
			UserID uint
			Promos []string `gorm:"type:text[]"`
		}
		return tx.Migrator().CreateTable(&userPromos{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("user_promos")
	},
}
