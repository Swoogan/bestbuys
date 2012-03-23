package domain

//import "fmt"

const Quantity = 10

type Purchase struct {
	Name     string
	Cost     Money
	Quantity int

	FinanceIn  Finance
	FinanceOut  Finance
	Hours     int
	Ratio     Money
	NewIncome Money
}

type FullPurchase struct {
	First    Purchase
	Second   Purchase
	Increase Money
}

func (s Structure) purchase(finance Finance) FullPurchase {
	var result FullPurchase
	quantity := s.quantityToPurchase(finance.Income)

	if s.BuiltOn.RetainAlways {
		p := purchase{finance, s.Name, s.Cost, s.Income, quantity}
		structure := calculate(p)

		p = purchase{structure.Finance, s.BuiltOn.Name, s.BuiltOn.Cost, s.BuiltOn.Income, quantity}
		land := calculate(p)

		result = FullPurchase{
			First:    structure,
			Second:   land,
			Increase: structure.Increase,
			Quantity: quantity,
		}
	} else {
		p := purchase{finance, s.BuiltOn.Name, s.BuiltOn.Cost, s.BuiltOn.Income, quantity}
		land := calculate(p)

		p = purchase{land.Finance, s.Name, s.Cost, s.Income, quantity}
		structure := calculate(p)

		result = FullPurchase{
			First:          land,
			Second:         structure,
			IncomeIncrease: structure.IncomeIncrease,
			Quantity:       quantity,
		}
	}

	return result
}

func calculate(p purchase) PurchaseResult {
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

func (s Structure) timeToPurchase(quantity int, income, cost Money) Money {
	landHours := Money(0)

	if !s.BuiltOn.RetainAlways {
		landHours = (s.BuiltOn.Cost * Money(quantity)) / income
		income += s.BuiltOn.Income * Money(quantity)
	}

	structureHours := (s.Cost * Money(quantity)) / income
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
