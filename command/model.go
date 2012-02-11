package main

//
// MODEL
//
type game struct {
	finance finance
	monies monies
}

type finance struct {
	income int64
	upkeep int64
}

func (f finance) difference() int64 {
	return f.income - f.upkeep
}

type monies struct {
	balance int64
	wallet int64
	lands int64
}

func (m monies) total() int64 {
	return m.balance + m.wallet + m.lands
}

