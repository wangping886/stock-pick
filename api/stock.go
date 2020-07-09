package api

import (
	"github.com/labstack/echo"
	"github.com/wangping886/stock-pick/dao"
	"log"
)

func StockInfo(ectx echo.Context) error {
	st, err := dao.SelectStockTrade("600521")
	if err != nil {
		log.Println("func", "SelectStockTrade", "code", "", "errmsg", err.Error())
		ectx.JSON(500, nil)
	}
	ectx.JSON(200, st)
	return nil
}
