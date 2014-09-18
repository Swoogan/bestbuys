package domain

import (
	"time"

        "gopkg.in/mgo.v2/bson"
)

type Game struct {
	Id         bson.ObjectId "hack" // this should be "_id" but then Upsert doesn't do anything
	Finance    Finance
	Structures map[string]Structure
	Purchases  Result
	LastSaved  time.Time
}

type Finance struct {
	Income Money
	Upkeep Money

	Balance Money
	Wallet  Money
	Lands   Money
}

func (f Finance) Hourly() Money {
	return f.Income - f.Upkeep
}

func (f Finance) Daily(hourly Money) Money {
	return hourly * CollectionTime
}

func (f Finance) TotalMoney() Money {
	return f.Balance + f.Wallet + f.Lands
}

type Land struct {
	Name         string
	Cost         Money
	Income       Money
	RetainAlways bool
}
