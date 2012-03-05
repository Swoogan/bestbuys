package main

import (
	"math"
	"domain"
	"launchpad.net/gobson/bson"
)

func (s domain.Structure) IncreasePrice(quantity money) {
	s.cost += s.increase * quantity
}

func (s domain.Structure) timeToPurchase(quantity, income, cost domain.Money) domain.Money {
	hoursToBase := domain.Money(0)

	if !s.builtOn.RetainAlways {
		hoursToBase = (s.builtOn.Cost * quantity) / income
		income += s.builtOn.Income * quantity
	}

	var hoursToLand = (s.cost * quantity) / income
	var total = hoursToBase + hoursToLand
	var adjustedTotal = money(math.Ceil(float64(hoursToBase)) + math.Ceil(float64(hoursToLand)))

	if total <= 1 {
		return 1
	}

	return adjustedTotal
}

func (s domain.Structure) quantityToPurchase(income domain.Money) domain.Money {
	count := 2
	cci := domain.Money(0)
	for count <= 10 {
		cci += (count - 1) * s.Increase
		totalHours := s.TimeToPurchase(count, income, s.Cost)
		potentialIncome := OpportunityCost(income, totalHours)
		if potentialIncome > cci {
			break
		}
		count++
	}
	return math.Ceil(count - 1, 1)
}

func (s domain.Structure) opportunityCost(income, totalHours domain.Money) domain.Money {
	hours := s.timeToPurchase(1, income, s.Cost)
	numberOfProperties := 2
	oppotunityCost := domain.Money(0)

	for hours < totalHours {
		income += s.Income
		n := numberOfProperties - 1
		delta := math.Ceil(totalHours - hours, 0)
		opportunityCost += delta * s.Income
		totalCost = s.Cost * (n * s.Increase)
		hours += s.timeToPurchase(1, income, totalCost)
		numberOfProperties++
	}
	return opportunityCost
}
