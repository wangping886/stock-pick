package model

import "time"

type StockPotential struct {
	Id          int
	Type        int
	Code        int
	Name        string
	TotalDrop   float64
	FirstLevel  float64
	SecondLevel float64
	ThirdLevel  float64
	FourthLevel float64
	FifthLevel  float64
	UptonowDrop float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
