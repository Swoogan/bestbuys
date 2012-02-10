package main

import (
	"http"
	"log"
	"launchpad.net/mgo"
	"github.com/Swoogan/rest.go"
)

func main() {
	log.Printf("Connecting to mongodb")

	session, err := mgo.Mongo("localhost")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB("command")

	dispatcher := newDispatcher()
	tr := newTaskRest(db.C("tasks"), dispatcher)
	rest.Resource("tasks", tr)

	log.Printf("About to listen on 4041")
	err = http.ListenAndServe(":4041", nil)
	if err != nil {
		log.Fatal(err)
	}
}
