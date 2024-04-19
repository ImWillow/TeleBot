package handlers

import (
	"context"
	"fmt"
	"strings"
	"telegrambot/models"
	"telegrambot/utils"
	"time"

	"github.com/go-telegram/bot"
	m "github.com/go-telegram/bot/models"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	RegisterUser(ctx context.Context, b *bot.Bot, update *m.Update)
}

type handler struct {
	membersFile string
}

func NewHandler(membersFile string) Handler {
	h := new(handler)
	h.membersFile = membersFile

	return h
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
	}
	if err := utils.AddUserToData(user); err != nil {
		logrus.Error(err)
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatId,
		Text:   fmt.Sprintf(models.AllowedNewMember, nickname),
	})
}

func (h *handler) DeleteAllChatMessages(ctx context.Context, b *bot.Bot, update *m.Update) {

}
