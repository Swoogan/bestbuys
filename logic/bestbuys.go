package main

import (
	"fmt"
	//  "flag"
	//  "launchpad.net/gobson/bson"
	"launchpad.net/mgo"
)

func main() {
	// get the flags
	session, err := mgo.Mongo("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("uken").C("lands")

	//  land := &Land { name: "Scorched Sands", cost: 10000, income: 1000000, retainAlways: true }

	land := Land{}
	result := c.Find(nil)
	err = result.One(&land)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Count())
	fmt.Println(land.Name)
	fmt.Println(land.Income)
	//fmt.Println(land.income)

	/*
	   structure := NewStructure("Wakity", 1000000000, 15000000, 115000, land)
	   fmt.Println(structure.name)
	   fmt.Println(structure.cost)
	   structure.IncreasePrice(10)
	   fmt.Println(structure.cost)
	*/
}

