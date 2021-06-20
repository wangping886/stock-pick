package logic

import (
	"fmt"
	"github.com/wangping886/stock-pick/dao"
	"github.com/wangping886/stock-pick/model"
)

const (
	FirstLevel  = 2
	SecondLevel = 4
	ThirdLevel  = 6
	FourthLevel = 9
	FifthLevel  = 12
)

func CalculateRise() (err error) {
	//
	potentialList, err := dao.GetStockPotentialList(-50)
	if err != nil {
		return err
	}

	for _, p := range potentialList {
		var sp model.StockPotential
		sp.Id = p.Id
		startTime := p.CreatedAt.AddDate(0, 0, 1).Format("2006-01-02")
		trends, err := dao.SelectDaysBeforeStock(p.Code, startTime)
		if err != nil {
			fmt.Println("SelectDaysBeforeStock", err.Error(), p.Code, p.Id)
			continue
		}
		var accu float64
		for i, t := range trends {
			accu += t.GrowthRate
			if i == FirstLevel-1 {
				sp.FirstLevel = accu
			}
			if i == SecondLevel-1 {
				sp.SecondLevel = accu
			}
			if i == ThirdLevel-1 {
				sp.ThirdLevel = accu
			}
			if i == FourthLevel-1 {
				sp.FourthLevel = accu
			}
			if i == FifthLevel-1 {
				sp.FifthLevel = accu
			}
		}
		sp.UptonowDrop = accu
		err = dao.UpdateStockPotential(sp)
		if err != nil {
			fmt.Println("UpdateStockPotential", err.Error(), sp)
			continue
		}
	}
	if err != nil {
		return err
	}
	return nil
}
