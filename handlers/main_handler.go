package handlers

import (
	"context"
	"fmt"
	"strings"
	"telegrambot/models"
	"telegrambot/repos"
	"time"

	"github.com/go-telegram/bot"
	m "github.com/go-telegram/bot/models"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	UpdateWelcomeMSG(welocomeMSG *m.Message)
	RegisterUser(ctx context.Context, b *bot.Bot, update *m.Update)
	WelcomeHandler(ctx context.Context, b *bot.Bot, update *m.Update)
	GetPromos(ctx context.Context, b *bot.Bot, update *m.Update)
	GetMembers(ctx context.Context, b *bot.Bot, update *m.Update)
	GetCommands(ctx context.Context, b *bot.Bot, update *m.Update)
}

type handler struct {
	repos       repos.Repos
	welocomeMSG *m.Message
}

func NewHandler(repos repos.Repos) Handler {
	h := new(handler)
	h.repos = repos

	return h
}

func (h *handler) UpdateWelcomeMSG(welocomeMSG *m.Message) {
	h.welocomeMSG = welocomeMSG
}

func (h *handler) RegisterUser(ctx context.Context, b *bot.Bot, update *m.Update) {
	chatId := update.Message.Chat.ID
	nickname, _ := strings.CutPrefix(update.Message.Text, models.Register)
	defer b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    chatId,
		MessageID: update.Message.ID,
	})
	if nickname == "" {
		msg, _ := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:              chatId,
			Text:                "Пожалуйста, введите в формате `/register {никнейм} без скобочек. Пример: /register Aldeshara",
			DisableNotification: true,
		})

		time.Sleep(time.Minute)

		b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    chatId,
			MessageID: msg.ID,
		})

		return
	}

	user := models.User{
		TelegramID: update.Message.From.Username,
		NickName:   nickname,
		Role:       models.Role_member,
	}
	if err := h.repos.UserRepo.NewUser(user); err != nil {
		logrus.Error(err)
		return
	}

	b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    chatId,
		MessageID: h.welocomeMSG.ID,
	})

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatId,
		Text:   fmt.Sprintf(models.AllowedNewMember, nickname),
	})
}

func (h *handler) WelcomeHandler(ctx context.Context, b *bot.Bot, update *m.Update) {
	if update.Message.NewChatMembers != nil {
		msg, _ := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:              update.Message.Chat.ID,
			Text:                models.NewMember,
			DisableNotification: true,
		})
		time.Sleep(time.Minute)
		h.welocomeMSG = msg
	}
}

func (h *handler) GetPromos(ctx context.Context, b *bot.Bot, update *m.Update) {
	promos, err := h.repos.PromoRepo.GetPromos()
	if err != nil {
		logrus.Error(err)
		return
	}

	text := ""
	for _, promo := range promos {
		text += fmt.Sprintf("\\<\\-\\-\\-\\-`%s`\\-\\-\\-\\-\\>\n>%s\n>Date: %s\n>Active: %t\n", promo.Key, promo.Reward, promo.Date, promo.Active)
	}

	if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      text,
		ParseMode: m.ParseModeMarkdown,
	}); err != nil {
		logrus.Debug(err)
	}
}

func (h *handler) GetMembers(ctx context.Context, b *bot.Bot, update *m.Update) {
	users, err := h.repos.UserRepo.GetUsers()
	if err != nil {
		logrus.Error(err)
		return
	}

	text := ""
	for _, user := range users {
		text += fmt.Sprintf("Telegram: @%s - Никнейм: %s; \n", user.TelegramID, user.NickName)
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:              update.Message.Chat.ID,
		Text:                text,
		DisableNotification: true,
		ReplyParameters: &m.ReplyParameters{
			MessageID: update.Message.ID,
			ChatID:    update.Message.Chat.ID,
		},
	})
}

func (h *handler) GetCommands(ctx context.Context, b *bot.Bot, update *m.Update) {
	commandList := "`/register` - регистрация нового пользователя; \n`/promos` - список актуальных промокодов; \n"
	logrus.Debug("ChatID: ", update.Message.Chat.ID)
	if update.Message.Chat.ID != models.Chat_ID {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:              update.Message.From.ID,
			Text:                commandList,
			DisableNotification: true,
		})
		if err != nil {
			logrus.Debug(err)
			return
		}

		return
	}

	msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:              update.Message.Chat.ID,
		Text:                commandList,
		DisableNotification: true,
	})
	if err != nil {
		logrus.Debug(err)
		return
	}

	go func() {
		time.Sleep(time.Minute)

		if _, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.Message.From.ID,
			MessageID: msg.ID,
		}); err != nil {
			logrus.Debug(err)
			return
		}
		if _, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    update.Message.From.ID,
			MessageID: update.Message.ID,
		}); err != nil {
			logrus.Debug(err)
			return
		}
	}()
}
