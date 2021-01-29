package dao

import (
	"context"
	"github.com/wangping886/stock-pick/db"
	"github.com/wangping886/stock-pick/model"
	"time"
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

func SelectStockTrade(code string) (model.StockTrend, error) {
	var st model.StockTrend
	query := "select id,name,code,opening_price,opening_growth_rate,closeing_price,yesterday_closing_price,growth_rate,stock_amplitude,volume,turnover_rate,market_value,trading_day,created_at,updated_at from stock_trend where code = ? "
	err := db.DB.QueryRowContext(context.TODO(), query, code).Scan(
		&st.Id,
		&st.Name,
		&st.Code,
		&st.OpeningPrice,
		&st.OpeningGrowthRate,
		&st.CloseingPrice,
		&st.YesterdayClosingPrice,
		&st.GrowthRate,
		&st.StockAmplitude,
		&st.Volume,
		&st.TurnoverRate,
		&st.MarketValue,
		&st.TradingDay,
		&st.CreatedAt,
		&st.UpdatedAt,
	)
	return st, err
}

func SelectStockCodes() ([]int, error) {
	var codes = make([]int, 0)
	query := "select distinct(code) from stock_trend  "
	rows, err := db.DB.Query(query)
	for rows.Next() {
		var code int
		rows.Scan(
			&code,
		)
		codes = append(codes, code)
	}
	return codes, err
}

func SelectDaysBeforeStock(code int, days int) ([]model.StockTrend, error) {
	var sts = make([]model.StockTrend, 0)
	beforeTime := time.Now().AddDate(0, 0, days).Format("2006-01-02")
	query := "select id,name,code,growth_rate,market_value from stock_trend where code = ? and trading_day >= ? "
	rows, err := db.DB.Query(query, code, beforeTime)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var st model.StockTrend
		rows.Scan(
			&st.Id,
			&st.Name,
			&st.Code,
			&st.GrowthRate,
			&st.MarketValue,
		)
		sts = append(sts, st)
	}
	return sts, nil
}
