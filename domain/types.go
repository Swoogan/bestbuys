package domain

import "launchpad.net/gobson/bson"

const CollectionTime = 24

type Money float64

type Data map[string]interface{}

type Event struct {
	Name string
	Date bson.Timestamp
	Data Data
}

type Structures map[string]Structure

type Result struct {
	Path  string
	Ratio Money
}

type PurchaseResult struct {
	Name      string
	Hours     int
	Quantity  int
	Cost      Money
	Ratio     Money
	NewIncome Money
}

type FullPurchase struct {
	First    PurchaseResult
	Second   PurchaseResult
	Increase Money
	Quantity int
}
