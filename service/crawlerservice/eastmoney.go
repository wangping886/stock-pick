package crawlerservice

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/wangping886/stock-pick/httpclient"
	"github.com/wangping886/stock-pick/model"
)

type EastCrawler struct {
	Url string
}

func NewEastCrawler(code int) *EastCrawler {
	c := new(EastCrawler)
	c.genUrl(code)
	return c
}

func (c *EastCrawler) genUrl(code int) {
	urlPrefix := "http://quote.eastmoney.com/sh"
	c.Url = fmt.Sprintf("%s%d%s", urlPrefix, code, ".html")
}
func (e *EastCrawler) ProcessTrendHtml() (data model.StockTrend, err error) {

	httpclient := httpclient.DefaultHttpclient()
	resp, err := httpclient.Get(e.Url)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(doc.Html())

	var headings, row []string
	var rows [][]string
	fmt.Println(doc.Find("#quote-digest").Html())
	doc.Find("#quote-digest").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
				headings = append(headings, tableheading.Text())
			})
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, tablecell.Text())
			})
			rows = append(rows, row)
			row = nil
		})
	})
	fmt.Println("####### headings = ", len(headings), headings)
	fmt.Println("####### rows = ", len(rows), rows)
	return data, nil
}

func (e *EastCrawler) ProcessNoticeHtml() (err error) {

	httpclient := httpclient.DefaultHttpclient()
	resp, err := httpclient.Get("http://data.eastmoney.com/notices/getdata.ashx?StockCode=&FirstNodeType=0&CodeType=1&PageIndex=1&PageSize=50&jsObj=ySLiMGsI&SecNodeType=0&Time=&rt=52618026")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(doc.Html())
	return
}
