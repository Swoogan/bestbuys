package domain

import "labix.org/v2/mgo/bson"

const CollectionTime = 24

type Money float64

type Data map[string]interface{}

type Event struct {
	Name string
	Date bson.MongoTimestamp
	Data Data
}

type Structures map[string]Structure

type Result struct {
	Path  string
	Ratio Money
}

