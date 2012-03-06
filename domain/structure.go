package domain

import (
	"math"
)

type Structure struct {
	Name     string
	Cost     Money
	Increase Money
	Income   Money
	BuiltOn *Land
}

func (s Structure) CalcualatePurchase(finance Finance) logic.FullPurchase {
	var result FullPurchase
	quantity := s.quantityToPurchase(finance.Income)

	if s.BuiltOn.RetainAlways {
		structure := s.CalcualatePurchase(finance, s.Name, s.Cost, s.Income, quantity)
		var land = CalcualatePurchase(structure.Finance, _builtOn.Name, _builtOn.Cost, _builtOn.Income, quantity);

		result = FullPurchase{structure, land, structure.IncomeIncrease, quantity}
	} else {
		land := CalcualatePurchase(finance, s.BuiltOn.Name, s.BuiltOn.Cost, s.BuiltOn.Income, quantity);
		structure := CalcualatePurchase(land.Finance, s.Name, s.Cost, s.Income, quantity);
		result = FullPurchase{land, structure, structure.IncomeIncrease, quantity}
    }

    return result
}


func (s Structure) IncreasePrice(quantity Money) {
	s.Cost += s.Increase * quantity
}

func (s Structure) timeToPurchase(quantity, income, cost Money) Money {
	landHours := Money(0)

	if !s.BuiltOn.RetainAlways {
		landHours = (s.BuiltOn.Cost * quantity) / income
		income += s.BuiltOn.Income * quantity
	}

	structureHours := (s.Cost * quantity) / income
	total := landHours + structureHours
	adjustedTotal := ceil(landHours) + ceil(structureHours)

	if total <= 1 {
		return 1
	}

	return adjustedTotal
}

func (s Structure) quantityToPurchase(income Money) Money {
	count := Money(2)
	cci := Money(0)
	for count <= 10 {
		cci += count - 1 * s.Increase
		totalHours := s.timeToPurchase(count, income, s.Cost)
		potentialIncome := s.opportunityCost(income, totalHours)
		if potentialIncome > cci {
			break
		}
		count++
	}
	return max(count - 1, 1)
}

func (s Structure) opportunityCost(income, totalHours Money) Money {
	hours := s.timeToPurchase(1, income, s.Cost)
	numberOfProperties := Money(2)
	opportunityCost := Money(0)

	for hours < totalHours {
		income += s.Income
		n := numberOfProperties - 1
		delta := max(totalHours - hours, 0)
		opportunityCost += delta * s.Income
		totalCost := s.Cost * (n * s.Increase)
		hours += s.timeToPurchase(1, income, totalCost)
		numberOfProperties++
	}
	return opportunityCost
}

func max(x Money, y float64) Money {
	return Money(math.Fmax(float64(x), y))
}

func ceil(x Money) Money {
	return Money(math.Ceil(float64(x)))
}
