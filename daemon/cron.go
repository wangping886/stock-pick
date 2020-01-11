package daemon

import (
	"github.com/wangping886/stock-pick/logic"
	"log"
)

func CrawlerStockCron() {

	err := logic.CrawlerTrendProcess()
	if err != nil {
		log.Println("logic.CrawlerProcess error:", err)
	}

}
