package crawlerservice

import (
	"fmt"
	"testing"
)

//func TestProcessHtml(t *testing.T) {
//
//	crawler := NewEastCrawler(600521)
//	data, err := crawler.ProcessTrendHtml()
//	fmt.Println(1, data.CloseingPrice, 2, err)
//}

func TestProcessNoticeHtml(t *testing.T) {

	crawler := NewEastCrawler(0)
	err := crawler.ProcessNoticeHtml()
	fmt.Println(1, 2, err)
}

//func TestExampleNewDocumentFromReader(t *testing.T) {
//	// create from a string
//	data := `
//<html>
//	<head>
//		<title>My document</title>
//	</head>
//
//	<body id="test">
//		<div id="div1">
//			<h1 id="2222">Header2</h1>
//			<h2 id="2222">Header22</h1>
//
//			<h1 id="3333">Header3</h1>
//			<h1 id="div1">Header4</h1>
//
//		</div>
//
//		<div id="div2">
//			<h1 id="2222">Header2222</h1>
//		</div>
//
//	</body>
//</html>`
//
//	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
//	if err != nil {
//		log.Fatal(err)
//	}
//	header := doc.Find("#div1").First().Children().First().Next().Text()
//	fmt.Printf("%+v", header)
//	goGet()
//	// Output: Header
//}
//
//func goGet() {
//	var headings, row []string
//	var rows [][]string
//	data := `<html>
// <head></head>
// <body>
//  <div class="content" id="zjlx_hqcont">
//   <table cellpadding="0" cellspacing="0" class="t2 ns2">
//    <thead>
//     <tr>
//      <th>代码</th>
//      <th>名称</th>
//      <th>最新价</th>
//      <th>涨跌额</th>
//      <th>涨跌幅(%)</th>
//      <th>振幅(%)</th>
//      <th>成交量(手)</th>
//      <th>成交额(万)</th>
//      <th>昨收</th>
//      <th>今开</th>
//      <th>最高</th>
//      <th>最低</th>
//      <th>换手率(%)</th>
//      <th>量比</th>
//      <th>市盈率</th>
//      <th>加自选</th>
//     </tr>
//    </thead>
//    <tbody>
//     <tr>
//      <td><a href="http://quote.eastmoney.com/600521.html" target="_blank">600521</a></td>
//      <td><a href="http://data.eastmoney.com/stockdata/600521.html" target="_blank">华海药业</a></td>
//      <td><span class="green">18.21</span></td>
//      <td><span class="green">-0.74</span></td>
//      <td><span class="green">-3.91</span></td>
//      <td><span>4.85</span></td>
//      <td><span>184188</span></td>
//      <td><span>34069.08</span></td>
//      <td>18.95</td>
//      <td><span class="red">19.10</span></td>
//      <td><span class="red">19.12</span></td>
//      <td><span class="green">18.20</span></td>
//      <td><span>1.47</span></td>
//      <td><span>0.80</span></td>
//      <td><span>35.29</span></td>
//      <td><a href="http://quote.eastmoney.com/favor/infavor.aspx?code=600521" title="将600521(华海药业)加为自选股"><img src="/images/add.gif" /></a></td>
//     </tr>
//    </tbody>
//   </table>
//  </div>
// </body>
//</html>`
//	//	data := `<html><body>
//	//<div id="zjlx_hqcont"">
//	//	<table id="tableid">
//	//		<tr><th>Heading 1</th><th>Heading two</th></tr>
//	//		<tr><td>Data 11</td><td>Data 12</td></tr>
//	//		<tr><td>Data 21</td><td>Data 22</td></tr>
//	//		<tr><td>Data 31</td><td>Data 32</td></tr>
//	//		<tr><td>Data 41</td><td>Data 42</td></tr>
//	//	</table>
//	//	<p>Stuff in here</p>
//	//	<table id="tableid">
//	//		<tr><th>Heading 21</th><th>Heading 2two</th></tr>
//	//		<tr><td>Data 211</td><td>Data 212</td></tr>
//	//		<tr><td>Data 221</td><td>Data 222</td></tr>
//	//		<tr><td>Data 231</td><td><span></span><span><a href="">Data 232</a></span></td></tr>
//	//		<tr><td>Data 241</td><td>Data 242</td></tr>
//	//	</table>
//	//</div>
//	//	</body>
//	//	</html>
//	//	`
//	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
//	if err != nil {
//		fmt.Println("No url found")
//		log.Fatal(err)
//	}
//
//	// Find each table
//	doc.Find("#zjlx_hqcont").Each(func(index int, tablehtml *goquery.Selection) {
//		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
//			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
//				headings = append(headings, tableheading.Text())
//			})
//			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
//				row = append(row, tablecell.Text())
//			})
//			rows = append(rows, row)
//			row = nil
//		})
//	})
//	fmt.Println("####### headings = ", len(headings), headings)
//	fmt.Println("####### rows = ", len(rows), rows)
//}
