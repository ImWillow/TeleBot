package main

import (
	"context"
	"os"
	"os/signal"
	"telegrambot/gorm"
	"telegrambot/handlers"
	"telegrambot/models"
	"telegrambot/promo"
	"telegrambot/repos"
	"time"

	"github.com/go-telegram/bot"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})

	// Create db conn
	gm := gorm.NewGormModule()
	if err := gm.Connect(); err != nil {
		log.WithField("error", err).Error("can't create db connection")
		return
	}
	if err := gm.AutoMigrate(); err != nil {
		log.WithField("error", err).Error("can't automigrate db")
		return
	}

	// Parse promos
	var done chan bool
	defer close(done)
	ticker := time.NewTicker(time.Hour * 24)
	go promo.StartParsing(done, ticker, gm.GetRM())
	defer func() {
		done <- true
	}()

	// Create repositories
	repos := repos.NewRepo(gm)

	// Create handler
	h := handlers.NewHandler(repos)

	// Bot implementation.
	opts := []bot.Option{
		bot.WithDefaultHandler(h.WelcomeHandler),
	}

	b, err := bot.New(models.BotAPI, opts...)
	if err != nil {
		log.WithField("error", err).Error("can't create bot")
		return
	}

	var doneSender chan bool
	defer close(done)
	tickerSender := time.NewTicker(time.Hour * 1)
	go func() {
		time.Sleep(time.Minute)
		promo.StartSendNewPromos(doneSender, tickerSender, gm.GetRM(), b)
	}()
	defer func() {
		done <- true
	}()

	b.RegisterHandler(bot.HandlerTypeMessageText, models.Register, bot.MatchTypePrefix, h.RegisterUser)
	b.RegisterHandler(bot.HandlerTypeMessageText, models.Members, bot.MatchTypePrefix, h.GetMembers)
	b.RegisterHandler(bot.HandlerTypeMessageText, models.Commands, bot.MatchTypePrefix, h.GetCommands)

	log.Debug("Start bot")
	b.Start(ctx)
}
