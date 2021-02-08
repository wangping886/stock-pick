package dao

import (
	"github.com/wangping886/stock-pick/db"
	"time"
)

type StockPotential struct {
	Id        int
	Type      int
	Code      int
	Name      string
	TotalDrop float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InsertStockPotential(sp StockPotential) error {
	sql := "insert into stock_potential(type,code,name,total_drop) values(?,?,?,?)"
	_, err := db.DB.Exec(sql, sp.Type, sp.Code, sp.Name, sp.TotalDrop)
	return err
}
