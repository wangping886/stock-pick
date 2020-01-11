package apiservice

import (
	"fmt"
	"testing"
)

func TestProcessHtml(t *testing.T) {

	api := NewJuheService("e09b5e826762843f25cb31bb13f7bdcf")
	data, err := api.GetStockInfo()
	fmt.Println(1, data, 2, err)

}
