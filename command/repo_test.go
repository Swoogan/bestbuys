package main

import (
	"log"

        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"

        "bitbucket.org/Swoogan/bestbuys/domain"
)

func StoreTest() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB("command")
	repo := newRepository()
	id := bson.NewObjectId()
	log.Println("New id is:", id.Hex())
	repo[id.Hex()] = domain.Game{
		Id:      id,
		Finance: domain.Finance{1, 0, 0, 1, 0},
	}
	log.Println("Finance income is:", repo[id.Hex()].Finance.Income)

	repo.snapshot(db)
}

func LoadTest() {
	session, err := mgo.Dial("localhost")
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
