package promo

import (
	"net/http"
	dbmodels "telegrambot/gorm/models"
	"telegrambot/gorm/requests"
	"telegrambot/models"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		logrus.Errorf("HTTP Error %d: %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.WithError(err).Error("can't create document from page")
	}

	var promos []dbmodels.Promo
	doc.Find("body > div.at-wrap > div.xxl_container > section > main > div > div > div.games-content > div.codes-module").
		Find("div").
		Each(func(i int, s *goquery.Selection) {
			if s.HasClass("item-promo module md-block") {
				var promo dbmodels.Promo
				promo.Key = s.Find("h4").Text()
				promo.Reward = s.Find("p").Text()
				promo.Date = s.Find(`div`).Nodes[0].FirstChild.Data
				if s.Find(`div`).Nodes[1].FirstChild.Data == "Активный" {
					promo.Active = true
				} else {
					promo.Active = false
				}
				promos = append(promos, promo)
			}
		})

	// Add promos to db
	if err := rm.ClearPromos(); err != nil {
		logrus.WithError(err).Error("can't delete promos from db")
	}
	for _, promo := range promos {
		if err := rm.NewPromo(promo); err != nil {
			logrus.WithError(err).Error("can't write promo to db")
		}
	}
}
