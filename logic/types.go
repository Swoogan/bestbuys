package logic

import "domain"

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
