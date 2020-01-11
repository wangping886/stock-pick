package logic

import (
	"github.com/wangping886/stock-pick/dao"
	"github.com/wangping886/stock-pick/service/apiservice"
	"log"
	"strconv"
	"time"
)

func CrawlerTrendProcess() error {
	now := time.Now()
	//上海
	for i := 1; i < 25; i++ {

		api := apiservice.NewJuheService("e09b5e826762843f25cb31bb13f7bdcf")
		api.ApiAddr = "http://web.juhe.cn:8080/finance/stock/shall"
		api.Page = strconv.Itoa(i)
		api.Type = "4"
		stocks, err := api.GetStockInfo()
		if err != nil {
			log.Println("sh api.GetStockInfo error:", err)
		}
		for _, v := range stocks {
			err := dao.InsertStockTrade(v)
			if err != nil {
				log.Println("dao.InsertStockTrade error:", err)
				continue
			}
		}
	}
	//深圳
	for i := 1; i < 50; i++ {

		api := apiservice.NewJuheService("e09b5e826762843f25cb31bb13f7bdcf")
		api.ApiAddr = "http://web.juhe.cn:8080/finance/stock/szall"
		api.Page = strconv.Itoa(i)
		api.Type = "4"
		stocks, err := api.GetStockInfo()
		if err != nil {
			log.Println("sz api.GetStockInfo error:", err)
		}
		for _, v := range stocks {
			err := dao.InsertStockTrade(v)
			if err != nil {
				log.Println("dao.InsertStockTrade error:", err)
				continue
			}
		}
	}
	log.Println("crawler exec", time.Now().Sub(now).Seconds())
	return nil
}
