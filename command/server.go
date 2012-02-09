package main

import (
	"http"
	"log"
	"launchpad.net/mgo"
	"github.com/Swoogan/rest.go"
)

type Data map[string]interface{}

type Command struct {
	name string
	data Data
}

func main() {
	log.Printf("Connecting to mongodb")

	session, err := mgo.Mongo("localhost")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB("command")

	tr := NewTaskRest(db.C("tasks"), handleCommand)
	rest.Resource("tasks", tr)

	log.Printf("About to listen on 4041")
	err = http.ListenAndServe(":4041", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleCommand(cmd Command) {
	log.Println("Command Handler")
	handlers[cmd.name].

}
