package main

import (
	"launchpad.net/gobson/bson"
)

type money float64

type game struct {
	Id        bson.ObjectId "hack" // this should be "_id" but then Upsert doesn't do anything
	Finance   finance
	Monies    monies
	Lands []land
	LastSaved bson.Timestamp
}

type finance struct {
	Income money
	Upkeep money
}

func (f finance) hourly() money {
	return f.Income - f.Upkeep
}

func (f finance) daily(hourly money) money {
	return hourly * 24
}

type monies struct {
	Balance money
	Wallet  money
	Lands   money
}

func (m monies) total() money {
	return m.Balance + m.Wallet + m.Lands
}

type land struct {
	Name string
	Cost money
	Income money
}
