package logic

import (
	"github.com/wangping886/stock-pick/dao"
	"log"
)

func FiterPotentialStock() error {

	codes, err := dao.SelectStockCodes()
	if err != nil {
		return err
	}

	for _, code := range codes {
		//8  6天的数据
		stocks, err := dao.SelectDaysBeforeStock(code, -7)
		if err != nil {
			continue
		}
		var accumu float64
		var have4 bool
		for _, s := range stocks {
			accumu += s.GrowthRate
		}
		if accumu > -10 || stocks[0].MarketValue < 700 {
			continue
		}

		for _, s := range stocks {
			if s.GrowthRate <= -4.5 {
				have4 = true
			}
		}

		if !have4 {
			continue
		}

		sp := dao.StockPotential{
			Type: 1,
			Code: code,
			Name: stocks[0].Name,
		}
		err = dao.InsertStockPotential(sp)
		if err != nil {
			log.Println("error", err, "data", sp)
			continue
		}

	}
	return nil
}
