package main

import (
	"context"
	"os"
	"os/signal"
	"telegrambot/gorm"
	"telegrambot/handlers"
	"telegrambot/models"
	"telegrambot/repos"
	"time"

	"github.com/go-telegram/bot"
	m "github.com/go-telegram/bot/models"
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
	// Create repositories
	repos := repos.NewRepo(gm)
	// Create handler
	h := handlers.NewHandler(repos)

	// Bot implementation.
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("6870464352:AAFGYIf7A2s3MamrEkiE93lZifvo_NfBa7w", opts...)
	if err != nil {
		log.WithField("error", err).Error("can't create bot")
		return
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, models.Register, bot.MatchTypePrefix, h.RegisterUser)

	log.Debug("Start bot")
	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *m.Update) {
	if update.Message.NewChatMembers != nil {
		msg, _ := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:              update.Message.Chat.ID,
			Text:                models.NewMember,
			DisableNotification: true,
		})
		time.Sleep(time.Minute)
		b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.Message.Chat.ID,
			MessageID: msg.ID,
		})
	}
}
