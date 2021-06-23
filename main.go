package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wangping886/stock-pick/api"
	"github.com/wangping886/stock-pick/daemon"
	"gopkg.in/robfig/cron.v2"
	"log"
	"net/http"
	"time"
)

func main() {
	// Echo instance
	engine := echo.New()

	// Middleware
	engine.Use(middleware.Recover())

	engine.HEAD("/", func(ctx echo.Context) error {
		ctx.Response().WriteHeader(200)
		return nil
	})
	api.AddRouter(engine)

	go engine.StartServer(&http.Server{
		Addr:              ":1818",
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 2,
	})
	cron := cron.New()
	cron.Start()
	defer cron.Stop()
	log.Println("start server: stock-pick")
	//抓取每天收盘价
	cron.AddFunc("0 57 15 * * *", daemon.CrawlerStockCron)
	//筛选符合规则的潜力股
	cron.AddFunc("0 59 15 * * *", daemon.FiterPotentialStock)
	//计算潜力股的累计涨幅
	cron.AddFunc("0 02 16 * * *", daemon.CalculatePeriodRise)

	<-(chan int)(nil) // TODO: 替换成监听系统信号
}
