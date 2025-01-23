package promo

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	dbmodels "telegrambot/gorm/models"
	"telegrambot/gorm/requests"
	"telegrambot/models"
	"telegrambot/utils"

	m "github.com/go-telegram/bot/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-telegram/bot"
	"github.com/sirupsen/logrus"
)

func StartParsing(done chan bool, ticker *time.Ticker, rm requests.RequestModels) {
	parse(rm)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			parse(rm)
		}
	}
}

func parse(rm requests.RequestModels) {
	logrus.Debug("Start parsing")
	res, err := http.Get(models.Promo_URL)
	if err != nil {
		logrus.WithError(err).Error("can't get page")
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		logrus.Errorf("HTTP Error %d: %s", res.StatusCode, res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.WithError(err).Error("can't create document from page")
		return
	}

	var promos []dbmodels.Promo
	doc.Find("body > div.at-wrap > div.xxl_container > section > main > div > div > div.games-content > div.codes-module").
		Each(func(index int, table *goquery.Selection) {
			table.Find("h4.ajax-copy-text").
				Each(func(rowIndex int, row *goquery.Selection) {
					rNode := row.Next()
					sNode := rNode.Next()
					dateStr := sNode.Text()
					dateStr = utils.ReplaceDate(dateStr)
					date, err := time.Parse("2 01 2006", dateStr)
					if err != nil {
						fmt.Printf("Ошибка при парсинге даты '%s': %v \n", dateStr, err)
						return
					}
					fmt.Println(time.Since(date))
					if time.Since(date) < time.Hour*2000 {
						var code dbmodels.Promo
						code.Key = strings.TrimSpace(row.Text())
						code.Reward = strings.TrimSpace(rNode.Text())
						promos = append(promos, code)
					}
				})
		})

	dbPromos, err := rm.GetPromos()
	if err != nil {
		logrus.WithError(err).Error("can't create document from page")
		return
	}

	rp := checkPromo(dbPromos, promos)
	for _, promo := range rp {
		logrus.Debug("Add new promocode")
		if err := rm.NewPromo(promo); err != nil {
			logrus.WithError(err).Error("can't write promo to db")
		}
	}
}

func checkPromo(oldPromos, newPromos []dbmodels.Promo) []dbmodels.Promo {
	var truePromos []dbmodels.Promo
	if len(oldPromos) == 0 {
		return newPromos
	}
	for _, newPromo := range newPromos {
		f := false
		for _, oldPromo := range oldPromos {
			if newPromo.Key == oldPromo.Key {
				f = true
				break
			}
		}
		if f {
			continue
		}
		truePromos = append(truePromos, newPromo)
	}

	return truePromos
}

func StartSendNewPromos(done chan bool, ticker *time.Ticker, rm requests.RequestModels, bot *bot.Bot) {
	send(rm, bot)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			send(rm, bot)
		}
	}
}

func send(rm requests.RequestModels, b *bot.Bot) {
	dbPromos, err := rm.GetPromos()
	if err != nil {
		logrus.WithError(err).Error("can't create document from page")
		return
	}
	for _, promo := range dbPromos {
		if !promo.Sended {
			if _, err := b.SendMessage(context.Background(), &bot.SendMessageParams{
				ChatID:              models.Chat_ID,
				Text:                fmt.Sprintf("\\#\\#`%s`\\#\\#\n>%s\n", promo.Key, promo.Reward),
				DisableNotification: true,
				ParseMode:           m.ParseModeMarkdown,
				MessageThreadID:     models.PromoId,
			}); err != nil {
				logrus.WithError(err).Error("can't send message to chat")
				continue
			}
			promo.Sended = true
			if err = rm.UpdatePromo(promo); err != nil {
				logrus.WithError(err).Error("can't update promo")
				continue
			}
		}
	}
}
