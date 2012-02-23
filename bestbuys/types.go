package bestbuys

import "launchpad.net/gobson/bson"

type Money float64

type Data map[string]interface{}

type Event struct {
	Name string
	Date bson.Timestamp
	Data Data
}
