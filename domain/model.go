package domain

import (
	"launchpad.net/gobson/bson"
)

type Game struct {
	Id         bson.ObjectId "hack" // this should be "_id" but then Upsert doesn't do anything
	Finance    Finance
	Monies     Monies
	Structures []Structure
	Purchases Result
	LastSaved  bson.Timestamp
}

type Finance struct {
	Income Money
	Upkeep Money
}

func (f Finance) Hourly() Money {
	return f.Income - f.Upkeep
}

func (f Finance) Daily(hourly Money) Money {
	return hourly * 24
}

type Monies struct {
	Balance Money
	Wallet  Money
	Lands   Money
}

func (m Monies) Total() Money {
	return m.Balance + m.Wallet + m.Lands
}

type Land struct {
	Name   string
	Cost   Money
	Income Money
	RetainAlways bool
}
