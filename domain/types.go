package domain

import "launchpad.net/gobson/bson"

const CollectionTime = 24

type Money float64

type Data map[string]interface{}

type Event struct {
	Name string
	Date bson.Timestamp
	Data Data
}

type Structures map[string]Structure

type Buy struct {
	Name string
	Price Money
	Hours int
	NewIncome Money
	IncomePerHour Money
}

func NewBuy(p *Purchase) Buy {
	return Buy { p.Name, p.TotalCost, p.Hours, p.FinanceOut.Income, p.Ratio }
}

type Result struct {
	Path  []Buy
	Ratio Money
}
