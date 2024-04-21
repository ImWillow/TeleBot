package migrations

import (
	"telegrambot/gorm/migrations/migs"

	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations []*gormigrate.Migration = []*gormigrate.Migration{
	migs.M1, migs.M2,
}
