package crawlerservice

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/wangping886/stock-pick/httpclient"
	"github.com/wangping886/stock-pick/model"
)

type SinaCrawler struct {
	Url string
}

func NewSinaCrawler(code int) *SinaCrawler {
	c := new(SinaCrawler)
	c.genUrl(code)
	return c
}

func (c *SinaCrawler) genUrl(code int) {
	urlPrefix := "http://finance.sina.com.cn/realstock/company/sh600521/nc.shtml"
	c.Url = fmt.Sprintf("%s%d%s", urlPrefix, code, "/nc..html")
}
func (s *SinaCrawler) ProcessHtml() (data model.StockTrend, err error) {

	httpclient := httpclient.DefaultHttpclient()
	resp, err := httpclient.Get(s.Url)
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
	fmt.Println(doc.Find("#hqDetails").Html())
	doc.Find("#hqDetails").Each(func(index int, tablehtml *goquery.Selection) {
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
