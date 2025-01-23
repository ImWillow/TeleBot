package utils

import (
	"strconv"
	"strings"
)

func StringsToInts(str []string) ([]int64, error) {
	ints := []int64{}
	for _, s := range str {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		ints = append(ints, int64(i))
	}

	return ints, nil
}

var months map[string]string = map[string]string{
	"января":   "01",
	"февраля":  "02",
	"марта":    "03",
	"апреля":   "04",
	"мая":      "05",
	"июня":     "06",
	"июля":     "07",
	"августа":  "08",
	"сентября": "09",
	"октября":  "10",
	"ноября":   "11",
	"декабря":  "12",
}

func ReplaceDate(s string) string {
	splited := strings.Split(s, " ")
	splited[1] = months[splited[1]]
	return strings.Join(splited, " ")
}
