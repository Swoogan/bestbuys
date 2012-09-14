package domain

import (
	"math"
)

const Quantity = 10

type Purchase struct {
	Name      string
	UnitCost  Money
	Income    Money
	FinanceIn Finance

	FinanceOut Finance
	TotalCost  Money
	Increase   Money
	Hours      int
	Ratio      Money
}

func NewPurchase(name string, cost Money, income Money, finance Finance) *Purchase {
	return &Purchase{Name: name, UnitCost: cost, Income: income, FinanceIn: finance}
}

func (p *Purchase) Calculate() {
	p.TotalCost = p.UnitCost * Quantity
	p.Increase = p.Income * Quantity

	p.FinanceOut = p.FinanceIn
	p.FinanceOut.Income += p.Increase

	if p.TotalCost <= p.FinanceIn.Wallet {
		p.FinanceOut.Wallet -= p.TotalCost
		p.Hours = 0
		p.Ratio = (p.Increase * p.FinanceIn.Income) / p.TotalCost
	} else {
		p.Hours = int(ceil(p.TotalCost / p.FinanceIn.Income))
		// Buy outright
		if p.Hours <= CollectionTime {
			moneyLeft := (Money(p.Hours) * p.FinanceIn.Income) - p.TotalCost
			p.FinanceOut.Balance += p.FinanceIn.Wallet * Money(0.9)
			p.FinanceOut.Wallet = moneyLeft
		} else { // Save to purchase
			dailyIncome := p.FinanceIn.Income * CollectionTime
			amountRemaining := p.TotalCost - dailyIncome
			withdrawl := min(p.FinanceIn.Balance, amountRemaining)
			remainderWithFee := (amountRemaining - withdrawl) * Money(1.1)

			p.TotalCost = dailyIncome + (amountRemaining * Money(1.1))
			p.Hours = CollectionTime + int(ceil(remainderWithFee/p.FinanceIn.Income))
			p.FinanceOut.Balance += p.FinanceIn.Wallet * Money(0.9)
			p.FinanceOut.Wallet = Money(0)
			p.FinanceOut.Balance -= withdrawl
		}

		p.Ratio = p.Increase / Money(p.Hours)
	}
}

func (p *Purchase) ShortName() string {
	if len(p.Name) < 25 {
		return p.Name
	}

	return p.Name[0:22] + "..."
}

//
// Helpers
//

func mmax(x Money, y float64) Money {
	return Money(math.Max(float64(x), y))
}

func imax(x int, y float64) int {
	return int(math.Max(float64(x), y))
}

func min(x Money, y Money) Money {
	return Money(math.Min(float64(x), float64(y)))
}

func ceil(x Money) Money {
	return Money(math.Ceil(float64(x)))
}
