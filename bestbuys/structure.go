package main

import (
  "math"
  "launchpad.net/gobson/bson"
)

type Land struct {
  game bson.ObjectId
  Name string
  Cost money
  Income money
  RetainAlways bool
}


type Structure struct {
  name string
  increase money
  builtOn *Land
  income money
  cost money
}

func NewStructure(name string, cost, increase, income money, land *Land) (*Structure) {
  return &Structure {
    name: name,
    cost: cost,
    increase: increase,
    income: income,
    builtOn: land,
  }
}

func (s *Structure) IncreasePrice(quantity money) {
  s.cost += s.increase * quantity
}

func (s *Structure) timeToPurchase(quantity money, income, cost money) (money) {
  hoursToBase := money(0)

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
