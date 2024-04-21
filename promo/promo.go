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
	doc.Find("#footable_20150 > tbody > tr").Each(func(i int, s *goquery.Selection) {
		var promo dbmodels.Promo
		promo.Key = s.Find("td").Nodes[0].FirstChild.Attr[1].Val
		promo.Reward = s.Find("td").Nodes[1].FirstChild.Data
		promos = append(promos, promo)
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
	// dbpromos, err := rm.GetPromos()
	// if err != nil {
	// 	logrus.WithError(err).Error("can't get promos from db")
	// }
	// for _, promo := range promos {
	// 	inner := true
	// 	for _, dbpromo := range dbpromos {
	// 		if dbpromo.Key == promo.Key {
	// 			inner = false
	// 			return
	// 		}
	// 	}
	// 	if !inner {
	// 		continue
	// 	}

	// 	if err := rm.NewPromo(promo); err != nil {
	// 		logrus.WithError(err).Error("can't write promo to db")
	// 	}
	// }
}
