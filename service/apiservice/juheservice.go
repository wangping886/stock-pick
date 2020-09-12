package apiservice

import (
	"encoding/json"
	"fmt"
	"github.com/wangping886/stock-pick/httpclient"
	"github.com/wangping886/stock-pick/model"
	"log"
	"math"
	"net/url"
	"strconv"
	"time"
)

type JuheService struct {
	ApiAddr string
	Appid   string
	Page    string
	Stock   string // a 股 b 股
	Type    string // 每页多少个
	Code    int
}

type JuheResponse struct {
	ErrorCode int `json:"error_code"`
	Result    struct {
		TotalCount string `json:"totalcount"`
		Page       string
		Num        string
		Data       []JuheStock
	}
}

type JuheDapanResponse struct {
	Resultcode string `json:"resultcode"`
	Reason     string `json:"reason"`
	Result     []JuheSingleStock
}

type JuheSingleStock struct {
	Data struct {
		Date          string
		Code          int
		IncrePer      string
		Name          string
		NowPri        string
		TodayStartPri string
		YestodEndPri  string
		TraAmount     string
	}
}
type JuheStock struct {
	Symbol        string
	Code          string
	Name          string
	Trade         string  //最新价
	//Pricechange   float64 // 涨跌额
	Changepercent interface{} //涨跌幅
	Settlement    string  //昨收
	Open          string  //开盘
	High          string  //最高
	Low           string  //最低
	Volume        int     //成交量
	Amount        int     //成交金额
	Ticktime      string  //数据时间
	Mktcap        float64 //市值
	Turnoverratio float64 //还手率
}

func (j *JuheService) genHttpUrl() string {
	return ""
}
func NewJuheService(appid string) *JuheService {
	return &JuheService{Appid: appid, Stock: "a", Type: "80"}
}
func (j *JuheService) GetStockInfo() ([]*model.StockTrend, error) {
	var (
		jr        JuheResponse
		st        []*model.StockTrend
		err       error
		resString string
	)

	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	v := url.Values{}
	v.Set("key", "e09b5e826762843f25cb31bb13f7bdcf")
	v.Set("page", j.Page)
	v.Set("stock", "a")
	v.Set("type", j.Type)
	neturl := fmt.Sprintf("%s?%s", j.ApiAddr, v.Encode())

	//url := "http://web.juhe.cn:8080/finance/stock/shall?key=e09b5e826762843f25cb31bb13f7bdcf&page=1&stock=a&type=20"
	client := httpclient.DefaultHttpclient()
	resString, err = client.HttpGet(neturl)
	if err != nil {
		return st, nil
	}
	err = json.Unmarshal([]byte(resString), &jr)
	if err != nil {
		return st, err
	}

	st = append(st)
	for _, r := range jr.Result.Data {
		codeInt, err := strconv.Atoi(r.Code)
		open, err := strconv.ParseFloat(r.Open, 64)
		close, err := strconv.ParseFloat(r.Trade, 64)
		var grothrate float64
		switch value:=r.Changepercent.(type) {
		case float64:
			grothrate=value
		case string:
			grothrate, _ = strconv.ParseFloat(value, 64)
		default:
			log.Println("uncaught type")
		}
		yesterdayClose, err := strconv.ParseFloat(r.Settlement, 64)
		openingGrowthRate := math.Round(((open-yesterdayClose)*100/yesterdayClose)*100) / 100
		high, err := strconv.ParseFloat(r.High, 64)
		low, err := strconv.ParseFloat(r.Low, 64)
		amplitude := math.Round(((high-low)*100/yesterdayClose)*100) / 100
		if err != nil {
			log.Println("convert data error:", err)
		}
		data := &model.StockTrend{
			Name:                  r.Name,
			Code:                  codeInt,
			OpeningPrice:          open,
			OpeningGrowthRate:     openingGrowthRate,
			CloseingPrice:         close,
			YesterdayClosingPrice: yesterdayClose,
			GrowthRate:            grothrate,
			StockAmplitude:        amplitude,
			Volume:                r.Amount / 10000,
			TurnoverRate:          math.Round(r.Turnoverratio*100) / 100,
			MarketValue:           int(r.Mktcap) / 10000,
			TradingDay:            time.Now().Format("2006-01-02"),
		}
		log.Println("stockinfo", r, "err", err)
		st = append(st, data)
	}
	return st, nil
}

func (j *JuheService) GetIndexStockInfo() (model.StockTrend, error) {

	var (
		jr        JuheDapanResponse
		st        model.StockTrend
		err       error
		resString string
	)

	v := url.Values{}
	v.Set("key", "e09b5e826762843f25cb31bb13f7bdcf")
	if j.Code == 399001 {
		v.Set("gid", "sz399001")
	}
	if j.Code == 399002 {
		v.Set("gid", "sh000001")
	}
	neturl := fmt.Sprintf("%s?%s", j.ApiAddr, v.Encode())

	client := httpclient.DefaultHttpclient()
	resString, err = client.HttpGet(neturl)
	if err != nil {
		return st, nil
	}
	err = json.Unmarshal([]byte(resString), &jr)
	if err != nil {
		return st, err
	}
	dapanData := jr.Result[0].Data
	open, err := strconv.ParseFloat(dapanData.TodayStartPri, 64)
	close, err := strconv.ParseFloat(dapanData.NowPri, 64)
	yesterdayClose, err := strconv.ParseFloat(dapanData.YestodEndPri, 64)
	openingGrowthRate := math.Round(((open-yesterdayClose)*100/yesterdayClose)*100) / 100
	change, err := strconv.ParseFloat(dapanData.IncrePer, 64)
	amount, err := strconv.ParseFloat(dapanData.TraAmount, 64)

	data := model.StockTrend{
		Name:                  dapanData.Name,
		Code:                  j.Code,
		OpeningPrice:          open,
		OpeningGrowthRate:     openingGrowthRate,
		CloseingPrice:         close,
		YesterdayClosingPrice: yesterdayClose,
		GrowthRate:            change,
		Volume:                int(amount / 100000000.0),
		TradingDay:            time.Now().Format("2006-01-02"),
	}
	log.Println("stockinfo", data, "err", err)
	return data, nil
}
