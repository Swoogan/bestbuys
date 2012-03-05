package logic

import (
	"math"
	"domain"
)

func (s domain.Structure) IncreasePrice(quantity domain.Money) {
	s.Cost += s.Increase * quantity
}

func (s domain.Structure) timeToPurchase(quantity, income, cost domain.Money) domain.Money {
	landHours := domain.Money(0)

	if !s.BuiltOn.RetainAlways {
		landHours = (s.BuiltOn.Cost * quantity) / income
		income += s.BuiltOn.Income * quantity
	}

	structureHours := (s.Cost * quantity) / income
	total := landHours + structureHours

	clHours := math.Ceil(float64(landHours))
	csHours := math.Ceil(float64(structureHours))
	adjustedTotal := domain.Money(clHours + csHours)

	if total <= 1 {
		return 1
	}

	return adjustedTotal
}

func (s domain.Structure) quantityToPurchase(income domain.Money) domain.Money {
	count := 2
	cci := domain.Money(0)
	for count <= 10 {
		cci += domain.Money(count - 1) * s.Increase
		totalHours := s.timeToPurchase(count, income, s.Cost)
		potentialIncome := s.opportunityCost(income, totalHours)
		if potentialIncome > cci {
			break
		}
		count++
	}
	return domain.Money(math.Fmax(float64(count - 1), 1))
}

func (s domain.Structure) opportunityCost(income, totalHours domain.Money) domain.Money {
	hours := s.timeToPurchase(1, income, s.Cost)
	numberOfProperties := domain.Money(2)
	opportunityCost := domain.Money(0)

	for hours < totalHours {
		income += s.Income
		n := numberOfProperties - 1
		delta := domain.Money(math.Fmax(float64(totalHours - hours), 0))
		opportunityCost += delta * s.Income
		totalCost := s.Cost * (n * s.Increase)
		hours += s.timeToPurchase(1, income, totalCost)
		numberOfProperties++
	}
	return opportunityCost
}
