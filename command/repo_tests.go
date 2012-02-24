package main

import (
	"os"
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
	logger := log.New(os.Stderr, "", log.LstdFlags)
	repo := newRepository(logger)
	id := bson.NewObjectId()
	log.Println("New id is:", id.Hex())
	repo.games[id.Hex()] = game{
		Id:      id,
		Finance: finance{1, 0},
		Monies:  monies{0, 1, 0},
	}
	log.Println("Finance income is:", repo.games[id.Hex()].Finance.Income)

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
	logger := log.New(os.Stderr, "", log.LstdFlags)
	repo := newRepository(logger)
	repo.rebuild(db)

	for id, game := range repo.games {
		log.Println("Id is:", id)
		log.Println("Game is:", game)
	}
}
