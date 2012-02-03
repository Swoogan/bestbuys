package main

import (
	"http"
	"log"
	"launchpad.net/mgo"
	"bitbucket.org/Swoogan/mongorest"
)

func main() {
	log.Printf("Connecting to mongodb")

	session, err := mgo.Mongo("localhost")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB("uken")

	mongorest.NewMongoRest(db, "games")

	log.Printf("About to listen on 4040")
	err = http.ListenAndServe(":4040", nil)
	if err != nil {
		log.Fatal(err)
	}
}
