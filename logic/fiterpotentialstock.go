package logic

import (
	"github.com/wangping886/stock-pick/dao"
	"github.com/wangping886/stock-pick/model"
	"log"
	"time"
)

const (
	//6个交易日
	daysbefore = -7
	//累计跌幅
	totaldrop = -8.5
	//市值
	marketvalue = 300
	//单日跌幅
	somedaydrop = -2.9
)

func FiterPotentialStock() error {

	codes, err := dao.SelectStockCodes()
	if err != nil {
		return err
	}

	for _, code := range codes {
		//近期累计跌幅
		var accumu float64
		startTime := time.Now().AddDate(0, 0, daysbefore).Format("2006-01-02")
		stocks, err := dao.SelectDaysBeforeStock(code, startTime)
		if err != nil || len(stocks) == 0 {
			continue
		}
		var somedayBreakDrop bool
		for _, s := range stocks {
			accumu += s.GrowthRate
		}
		sp := model.StockPotential{
			Type:      1,
			Code:      code,
			Name:      stocks[0].Name,
			TotalDrop: accumu,
		}

		if code == 399001 {
			dao.InsertStockPotential(sp)
			continue
		}

		if accumu > totaldrop || stocks[0].MarketValue < marketvalue {
			continue
		}

		for _, s := range stocks {
			if s.GrowthRate <= -somedaydrop {
				somedayBreakDrop = true
			}
		}

		if !somedayBreakDrop {
			continue
		}
		err = dao.InsertStockPotential(sp)
		if err != nil {
			log.Println("error", err, "data", sp)
			continue
		}

	}
	return nil
}
