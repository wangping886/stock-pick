package daemon

import (
	"github.com/wangping886/stock-pick/logic"
	"log"
	"time"
)

var Suspension = map[string]bool{"2020-01-01": true, "2020-04-06": true, "2020-05-01": true, "2020-05-04": true, "2020-05-05": true,
	"2020-06-25": true, "2020-06-26": true,
	"2020-10-01": true, "2020-10-02": true, "2020-10-05": true, "2020-10-06": true, "2020-10-07": true, "2020-10-08": true,
	"2021-01-01": true, "2021-02-11": true, "2021-02-12": true, "2021-02-15": true, "2021-02-16": true, "2021-02-17": true,
	"2021-04-05": true, "2021-05-03": true, "2021-05-04": true, "2021-05-05": true, "2021-06-14": true, "2021-09-20": true,
	"2021-09-21": true, "2021-10-04": true, "2021-10-05": true, "2021-10-06": true, "2021-10-07": true,
}

func CrawlerStockCron() {
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday || Suspension[time.Now().Format("2006-01-02")] {
		log.Println("info:", "非交易时间")
		return
	}
	err := logic.CrawlerTrendProcess()
	if err != nil {
		log.Println("logic.CrawlerProcess error:", err)
	}

}

func FiterPotentialStock() {
	if time.Now().Weekday() == time.Saturday || time.Now().Weekday() == time.Sunday || Suspension[time.Now().Format("2006-01-02")] {
		log.Println("info:", "非交易时间")
		return
	}

	err := logic.FiterPotentialStock()
	if err != nil {
		log.Println("logic.FiterPotentialStock error:", err)
	}

}
