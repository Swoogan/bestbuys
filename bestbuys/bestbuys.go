package main

import (
  "fmt"
//  "flag"
)

func main() {
  // get the flags
  
  fmt.Println("Hello")
}


type Land struct {
  Name string
  Cost float64
  Income float64
  RetainAlways bool
}

type Result struct {
  Path string
  Ratio float64
}
