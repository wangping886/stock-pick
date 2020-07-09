package apiservice

import (
	"fmt"
	"testing"
)

//func TestProcessHtml(t *testing.T) {
//
//	api := NewJuheService("e09b5e826762843f25cb31bb13f7bdcf")
//	data, err := api.GetStockInfo()
//	fmt.Println(1, data, 2, err)
//
//}

func TestDapan(t *testing.T) {

	api := NewJuheService("e09b5e826762843f25cb31bb13f7bdcf")
	api.ApiAddr = "http://web.juhe.cn:8080/finance/stock/hs"
	api.Code = 399002
	data, err := api.GetIndexStockInfo()
	fmt.Println(1, data, 2, err)

}
