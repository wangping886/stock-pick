package main

import (
	"github.com/wangping886/stock-pick/daemon"
	"gopkg.in/robfig/cron.v2"
	"log"
)

const hushen = 600000

func main() {

	cron := cron.New()
	cron.Start()
	defer cron.Stop()
	log.Println("start server: stock-pick")
	cron.AddFunc("0 10 16 * * *", daemon.CrawlerStockCron)

	<-(chan int)(nil) // TODO: 替换成监听系统信号
}
