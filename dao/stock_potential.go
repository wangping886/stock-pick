package dao

import (
	"github.com/wangping886/stock-pick/db"
	"github.com/wangping886/stock-pick/model"
	"time"
)

func InsertStockPotential(sp model.StockPotential) error {
	sql := "insert into stock_potential(type,code,name,total_drop) values(?,?,?,?)"
	_, err := db.DB.Exec(sql, sp.Type, sp.Code, sp.Name, sp.TotalDrop)
	return err
}

func GetStockPotentialList(days int) ([]model.StockPotential, error) {
	var sps = make([]model.StockPotential, 0)
	beforeTime := time.Now().AddDate(0, 0, days).Format("2006-01-02")
	query := "select id,name,code,created_at from stock_potential where  trading_day >= ? order by id asc "
	rows, err := db.DB.Query(query, beforeTime)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var sp model.StockPotential
		rows.Scan(
			&sp.Id,
			&sp.Name,
			&sp.Code,
			&sp.CreatedAt,
		)
		sps = append(sps, sp)
	}
	return sps, nil
}

func UpdateStockPotential(sp model.StockPotential) (err error) {
	cmd := "update stock_potential set first_level= ? ,second_level = ? , third_level = ? , fourth_level = ? , fifth_level = ? , uptonow_drop = ? where id = ?"
	_, err = db.DB.Exec(cmd, sp.FirstLevel, sp.SecondLevel, sp.ThirdLevel, sp.FourthLevel, sp.FifthLevel, sp.UptonowDrop, sp.Id)
	return err
}
