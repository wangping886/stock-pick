package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func init() {
	var err error
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	DB, err = sql.Open("mysql", "root:root@tcp(000000.Mysql:3306)/stock_pick?clientFoundRows=false&parseTime=true&loc=Asia%2FShanghai&timeout=10s&collation=utf8mb4_bin&interpolateParams=true")
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(10)
	if err = DB.Ping(); err != nil {
		panic(err)
	}
}
