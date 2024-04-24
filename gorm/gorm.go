package gorm

import (
	"telegrambot/gorm/migrations"
	"telegrambot/gorm/requests"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormModule interface {
	Connect() error
	AutoMigrate() error
	GetRM() requests.RequestModels
}

type gormModule struct {
	db *gorm.DB
	rm requests.RequestModels
}

func NewGormModule() GormModule {
	gm := new(gormModule)

	return gm
}

func (gm *gormModule) Connect() error {
	dsn := "host=localhost user=postgres password=admin dbname=telebot port=5432 sslmode=disable TimeZone=Europe/Moscow"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	gm.db = conn
	gm.rm = requests.NewRequestModel(conn)

	return nil
}

func (gm *gormModule) AutoMigrate() error {
	m := gormigrate.New(gm.db, gormigrate.DefaultOptions, migrations.Migrations)
	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}

func (gm *gormModule) GetRM() requests.RequestModels {
	return gm.rm
}
