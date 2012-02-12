package event

import "launchpad.net/gobson/bson"

type Data map[string]interface{}

type Event struct {
	Name string
	Date bson.Timestamp
	Data Data
}
