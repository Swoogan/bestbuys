package main

import (
	"bestbuys"
	"launchpad.net/gobson/bson"
)

type game struct {
	Id         bson.ObjectId "hack" // this should be "_id" but then Upsert doesn't do anything
	Finance    finance
	Monies     monies
	Lands      []land
	Structures []structure
	LastSaved  bson.Timestamp
}

type finance struct {
	Income bestbuys.Money
	Upkeep bestbuys.Money
}

func (f finance) hourly() bestbuys.Money {
	return f.Income - f.Upkeep
}

func (f finance) daily(hourly bestbuys.Money) bestbuys.Money {
	return hourly * 24
}

type monies struct {
	Balance bestbuys.Money
	Wallet  bestbuys.Money
	Lands   bestbuys.Money
}

func (m monies) total() bestbuys.Money {
	return m.Balance + m.Wallet + m.Lands
}

type land struct {
	Name   string
	Cost   bestbuys.Money
	Income bestbuys.Money
}

type structure struct {
	Name     string
	Cost     bestbuys.Money
	Increase bestbuys.Money
	Income   bestbuys.Money
}
