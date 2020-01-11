package crawlerservice

import (
	"github.com/wangping886/stock-pick/model"
	"net/http"
)

type Crawler interface {
	ProcessHtml(res *http.Response) (model.StockTrend, error)
}
