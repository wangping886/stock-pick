package model

import "time"

type StockTrend struct {
	Id                    int64
	Name                  string
	Code                  int
	OpeningPrice          float64
	OpeningGrowthRate     float64
	CloseingPrice         float64
	YesterdayClosingPrice float64
	GrowthRate            float64
	StockAmplitude        float64
	Volume                int
	TurnoverRate          float64
	MarketValue           int
	TradingDay            string
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
