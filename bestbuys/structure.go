package main

import "math"

type Land struct {
  name string
  cost float64
  income float64
  retainAlways bool
}


type Structure struct {
  name string
  increase float64
  builtOn *Land
  income float64
  cost float64
}

func NewStructure(name string, cost, increase, income float64, land *Land) (*Structure) {
  return &Structure { 
    name: name, 
    cost: cost, 
    increase: increase, 
    income: income, 
    builtOn: land,
  }
}

func (s *Structure) IncreasePrice(quantity float64) {
  s.cost += s.increase * quantity
}

func (s *Structure) timeToPurchase(quantity float64, income, cost float64) (float64) {
  hoursToBase := float64(0)
  
  if !s.builtOn.retainAlways {
    hoursToBase = (s.builtOn.cost * quantity) / income
    income += s.builtOn.income * quantity
  }

  var hoursToLand = (s.cost * quantity) / income
  var total = hoursToBase + hoursToLand
  var adjustedTotal = math.Ceil(hoursToBase) + math.Ceil(hoursToLand)

  if total <= 1 {
    return 1
  }

  return adjustedTotal
}
