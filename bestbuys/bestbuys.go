package main

import (
  "fmt"
//  "flag"
)

func main() {
  // get the flags
  
  land := &Land { name: "Scorched Sands", cost: 10000, income: 1000000, retainAlways: true }
   
  fmt.Println(land.name)
  fmt.Println(land.income)

  structure := NewStructure("Wakity", 1000000000, 15000000, 115000, land)
  fmt.Println(structure.name)
  fmt.Println(structure.cost)
  structure.IncreasePrice(10)
  fmt.Println(structure.cost)
}


type Result struct {
  path string
  ratio uint64
}
