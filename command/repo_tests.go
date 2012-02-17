package main

import (
	"log"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
)

func StoreTest() {
	session, err := mgo.Mongo("localhost")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB("command")

	repo := newRepository()
	id := bson.NewObjectId()
	log.Println("New id is:", id.Hex())
	repo[id.Hex()] = game{
		Id:      id,
		Finance: finance{1, 0},
		Monies:  monies{0, 1, 0},
	}
	log.Println("Finance income is:", repo[id.Hex()].Finance.Income)

	repo.snapshot(db)
}

func LoadTest() {
	session, err := mgo.Mongo("localhost")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB("command")

	repo := newRepository()
	repo.rebuild(db)

	for id, game := range repo {
		log.Println("Id is:", id)
		log.Println("Game is:", game)
	}
}
