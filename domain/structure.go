package domain

import (
	"fmt"
	"math"
)

type Structure struct {
	Name     string
	Cost     Money
	Increase Money
	Income   Money
	BuiltOn  *Land
}

func (s *Structure) IncreasePrice() {
	s.Cost += s.Increase * Quantity
}

func (s *Structure) String() string {
	format := "%v\t\t %f\t %f\t %f\t"
	return fmt.Sprintf(format, s.Name, s.Cost, s.Increase, s.Income)
}

