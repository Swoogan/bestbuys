package domain

import "launchpad.net/gobson/bson"

type Money float64

type Data map[string]interface{}

type Event struct {
	Name string
	Date bson.Timestamp
	Data Data
}

type Result struct {
	path  string
	ratio uint64
}

type PurchaseResult struct {
    Name string
    Hours int
    Quantity int
    Cost domain.Money
    Ratio domain.Money
    Finance domain.Finance
    IncomeIncrease domain.Money
}

type FullPurchase struct {
    First PurchaseResult
    Second PurchaseResult
    IncomeIncrease domain.Money
    Quantity int
    Finance domain.Finance
    TotalHours int
}
