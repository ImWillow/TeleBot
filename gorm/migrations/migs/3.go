package migs

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

var M3 *gormigrate.Migration = &gormigrate.Migration{
	ID: "3",
	Migrate: func(tx *gorm.DB) error {
		type userPromos struct {
			gorm.Model
			UserID   uint
			PromoIDs pq.Int64Array `gorm:"type:integer[]"`
		}
		return tx.Migrator().CreateTable(&userPromos{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("user_promos")
	},
}
