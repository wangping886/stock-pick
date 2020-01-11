package dao

import (
	"context"
	"github.com/wangping886/stock-pick/db"
	"github.com/wangping886/stock-pick/model"
)

func InsertStockTrade(data *model.StockTrend) error {
	cmd := "insert into stock_trend (name,code,opening_price,opening_growth_rate,closeing_price,yesterday_closing_price,growth_rate,stock_amplitude,volume,turnover_rate,market_value,trading_day) values (?,?,?,?,?,?,?,?,?,?,?,?) "
	_, err := db.DB.ExecContext(context.TODO(), cmd,
		data.Name,
		data.Code,
		data.OpeningPrice,
		data.OpeningGrowthRate,
		data.CloseingPrice,
		data.YesterdayClosingPrice,
		data.GrowthRate,
		data.StockAmplitude,
		data.Volume,
		data.TurnoverRate,
		data.MarketValue,
		data.TradingDay,
	)
	return err
}
