package apiservice

import "github.com/wangping886/stock-pick/model"

type ApiService interface {
	GetStockInfo() model.StockTrend
}
