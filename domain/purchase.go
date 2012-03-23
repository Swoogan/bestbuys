package domain

import(
	"math"
)

const Quantity = 10

type Purchase struct {
	Name     string
	Cost     Money
	Income Money

	FinanceIn  Finance
	FinanceOut  Finance
	Increase     Money
	Hours     int
	Ratio     Money
}

func NewPurchase(name string, cost Money, income Money, finance Finance) *Purchase {
	return &Purchase{name, cost, income, finance}
}

func (p *Purchase) Calculate() {
	result := PurchaseResult{
		Name:     p.name,
		Quantity: p.quantity,
		Cost:     Money(p.quantity) * p.cost,
		Increase: p.income * Money(p.quantity),
	}

	result.Income += result.Increase

	if result.Cost <= p.finance.Wallet {
		result.Finance.Wallet -= result.Cost
		result.Hours = 0
		result.Ratio = (result.IncomeIncrease * p.finance.Income) / result.Cost
	} else {
		result.Hours = int(ceil(result.Cost / p.finance.Income))
		// Buy outright
		if result.Hours <= CollectionTime {
			moneyLeft := (Money(result.Hours) * p.finance.Income) - result.Cost
			result.Finance.Balance += p.finance.Wallet * Money(0.9)
			result.Finance.Wallet = moneyLeft
		} else { // Save to purchase
			dailyIncome := p.finance.Income * CollectionTime
			amountRemaining := result.Cost - dailyIncome
			withdrawl := min(p.finance.Balance, amountRemaining)
			remainderWithFee := (amountRemaining - withdrawl) * Money(1.1)

			result.Cost = dailyIncome + (amountRemaining * Money(1.1))
			result.Hours = CollectionTime + int(ceil(remainderWithFee/p.finance.Income))
			result.Finance.Balance += p.finance.Wallet * Money(0.9)
			result.Finance.Wallet = Money(0)
			result.Finance.Balance -= withdrawl
		}

		result.Ratio = result.IncomeIncrease / Money(result.Hours)
	}

	return result
}

func (p Purchase) Duration(s Structure) Money {
	landHours := Money(0)

	if !s.BuiltOn.RetainAlways {
		landHours = (s.BuiltOn.Cost * Quantity) / p.Income
		income += s.BuiltOn.Income * Quantity
	}

	structureHours := (s.Cost * Quantity) / p.Income
	total := landHours + structureHours
	adjustedTotal := ceil(landHours) + ceil(structureHours)

	if total <= 1 {
		return 1
	}

	return adjustedTotal
}


//
// Helpers
//

func mmax(x Money, y float64) Money {
	return Money(math.Fmax(float64(x), y))
}

func imax(x int, y float64) int {
	return int(math.Fmax(float64(x), y))
}

func min(x Money, y Money) Money {
	return Money(math.Fmin(float64(x), float64(y)))
}

func ceil(x Money) Money {
	return Money(math.Ceil(float64(x)))
}
