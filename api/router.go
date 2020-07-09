package api

import (
	"github.com/labstack/echo"
)

//AddRouter 增加路由
func AddRouter(e *echo.Echo) {
	routeGroup := e.Group("/stock-pick")
	routeGroup.GET("/stock-info", StockInfo)

}
