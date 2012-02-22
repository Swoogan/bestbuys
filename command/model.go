package main

import (
	"launchpad.net/gobson/bson"
)

type game struct {
	Id        bson.ObjectId "hack" // this should be "_id" but then Upsert doesn't do anything
	Finance   finance
	Monies    monies
	Lands []land
	LastSaved bson.Timestamp
}

type finance struct {
	Income int64
	Upkeep int64
}

func (f finance) hourly() int64 {
	return f.Income - f.Upkeep
}

func (f finance) daily(hourly int64) int64 {
	return hourly * 24
}

type monies struct {
	Balance int64
	Wallet  int64
	Lands   int64
}

func (m monies) total() int64 {
	return m.Balance + m.Wallet + m.Lands
}

type land struct {
	Name string
	Cost int64
	Income int64
}
