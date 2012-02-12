package main

type repository map[string]game

func newRepository() repository {
	return make(map[string]game, 3)
}

type game struct {
	finance finance
	monies  monies
}

type finance struct {
	income int64
	upkeep int64
}

func (f finance) hourly() int64 {
	return f.income - f.upkeep
}

func (f finance) daily(hourly int64) int64 {
	return hourly * 24
}

type monies struct {
	balance int64
	wallet  int64
	lands   int64
}

func (m monies) total() int64 {
	return m.balance + m.wallet + m.lands
}
