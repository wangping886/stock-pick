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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InsertStockPotential(sp StockPotential) error {
	sql := "insert into stock_potential(type,code,name) values(?,?,?)"
	_, err := db.DB.Exec(sql, sp.Type, sp.Code, sp.Name)
	return err
}
