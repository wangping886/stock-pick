package logic

import (
	"github.com/wangping886/stock-pick/dao"
)

func FiterPotentialStock() error {

	_, err := dao.SelectStockCodes()
	if err != nil {
		return err
	}
	return nil
}
