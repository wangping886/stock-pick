package daemon

import (
	"github.com/wangping886/stock-pick/logic"
	"log"
	"time"
)

var Suspension = map[string]bool{"2020-01-01": true, "2020-04-06": true, "2020-05-01": true, "2020-05-04": true, "2020-05-05": true,
	"2020-06-25": true, "2020-06-26": true,
	"2020-10-01": true, "2020-10-02": true, "2020-10-05": true, "2020-10-06": true, "2020-10-07": true, "2020-10-08": true}

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
