package main

type Structure struct {
  name string
  increase float64
  builtOn Land
  income float64
  cost float64
}

func NewStructure(name string, cost, increase, income float64, land Land) (*Structure) {
  return &Structure { 
    name: name, 
    cost: cost, 
    increase: increase, 
    income: income, 
    builtOn: land,
  }
}
