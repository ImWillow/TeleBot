package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"
	"telegrambot/models"
	"telegrambot/utils"

	"github.com/go-telegram/bot"
	m "github.com/go-telegram/bot/models"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	ShowMessageWithUserID(ctx context.Context, b *bot.Bot, update *m.Update)
}

type handler struct {
	membersFile string
}

func NewHandler(membersFile string) Handler {
	h := new(handler)
	h.membersFile = membersFile

	return h
}

func (h *handler) ShowMessageWithUserID(ctx context.Context, b *bot.Bot, update *m.Update) {
	chatId := update.Message.Chat.ID
	nickname, f := strings.CutPrefix(update.Message.Text, models.Register)
	if !f {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatId,
			Text:   "Пожалуйста, введите в формате `/register {никнейм} без скобочек",
		})
		return
	}
	inData, err := utils.CheckUserInData(nickname, h.membersFile)
	if err != nil {
		logrus.Debug(err)
		return
	}
	if !inData {
		file, err := os.OpenFile(h.membersFile, os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			logrus.Debug(err)
			return
		}
		if _, err := file.WriteString(fmt.Sprintf(models.UserTemplate, update.Message.From.Username, nickname)); err != nil {
			logrus.Debug(err)
			return
		}
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatId,
			Text:   fmt.Sprintf(models.AllowedNewMember, nickname),
		})
	} else {
		logrus.Debug("Пользователь уже зарегестрирован")
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.From.ID,
			Text:   fmt.Sprintf(models.AllowedNewMember, nickname),
		})
		b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    chatId,
			MessageID: update.Message.ID,
		})
	}
}
