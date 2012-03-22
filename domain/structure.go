package domain

import (
	"fmt"
	"math"
)

type purchase struct {
	finance  Finance
	monies   Monies
	name     string
	cost     Money
	income   Money
	quantity int
}

type Structure struct {
	Name     string
	Cost     Money
	Increase Money
	Income   Money
	BuiltOn  *Land
}

func (s Structure) purchase(finance Finance, monies Monies) FullPurchase {
	var result FullPurchase
	quantity := s.quantityToPurchase(finance.Income)

	if s.BuiltOn.RetainAlways {
		p := purchase{finance, monies, s.Name, s.Cost, s.Income, quantity}
		structure := calculate(p)

		p = purchase{structure.Finance, structure.Monies, s.BuiltOn.Name, s.BuiltOn.Cost, s.BuiltOn.Income, quantity}
		land := calculate(p)

		result = FullPurchase{
			First:          structure,
			Second:         land,
			IncomeIncrease: structure.IncomeIncrease,
			Quantity:       quantity,
		}
	} else {
		fmt.Println("Cost", s.Cost)
		fmt.Println("Name", s.Name)
		p := purchase{finance, monies, s.BuiltOn.Name, s.BuiltOn.Cost, s.BuiltOn.Income, quantity}
		land := calculate(p)

		p = purchase{land.Finance, land.Monies, s.Name, s.Cost, s.Income, quantity}
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
		Finance:        p.finance,
		Monies:         p.monies,
		Name:           p.name,
		Quantity:       p.quantity,
		Cost:           Money(p.quantity) * p.cost,
		IncomeIncrease: p.income * Money(p.quantity),
	}

	result.Finance.Income += result.IncomeIncrease

	if result.Cost <= p.monies.Wallet {
		result.Monies.Wallet -= result.Cost
		result.Hours = 0
		result.Ratio = (result.IncomeIncrease * p.finance.Income) / result.Cost
	} else {
		result.Hours = int(ceil(result.Cost / p.finance.Income))
		// Buy outright
		if result.Hours <= CollectionTime {
			moneyLeft := (Money(result.Hours) * p.finance.Income) - result.Cost
			result.Monies.Balance += p.monies.Wallet * Money(0.9)
			result.Monies.Wallet = moneyLeft
		} else { // Save to purchase
			dailyIncome := p.finance.Income * CollectionTime
			amountRemaining := result.Cost - dailyIncome
			withdrawl := min(p.monies.Balance, amountRemaining)
			remainderWithFee := (amountRemaining - withdrawl) * Money(1.1)

			result.Cost = dailyIncome + (amountRemaining * Money(1.1))
			result.Hours = CollectionTime + int(ceil(remainderWithFee/p.finance.Income))
			result.Monies.Balance += p.monies.Wallet * Money(0.9)
			result.Monies.Wallet = Money(0)
			result.Monies.Balance -= withdrawl
		}

		result.Ratio = result.IncomeIncrease / Money(result.Hours)
	}

	return result
}

func (s Structure) increasePrice(quantity int) {
	s.Cost += s.Increase * Money(quantity)
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

func (s Structure) quantityToPurchase(income Money) int {
	count := 2
	cci := Money(0)
	for count <= 10 {
		cci += Money(count-1) * s.Increase
		totalHours := s.timeToPurchase(count, income, s.Cost)
		potentialIncome := s.opportunityCost(income, totalHours)
		if potentialIncome > cci {
			break
		}
		count++
	}
	return imax(count-1, 1)
}

func (s Structure) opportunityCost(income, totalHours Money) Money {
	hours := s.timeToPurchase(1, income, s.Cost)
	numberOfProperties := Money(2)
	opportunityCost := Money(0)

	for hours < totalHours {
		income += s.Income
		n := numberOfProperties - 1
		delta := mmax(totalHours-hours, 0)
		opportunityCost += delta * s.Income
		totalCost := s.Cost * (n * s.Increase)
		hours += s.timeToPurchase(1, income, totalCost)
		numberOfProperties++
	}
	return opportunityCost
}

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
