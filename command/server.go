package main

import (
	"os"
	"log"
	"fmt"
	"http"
	"syscall"
	"os/signal"
	"launchpad.net/mgo"
	"github.com/Swoogan/rest.go"
)

/*
func main() {
	StoreTest()
	LoadTest()
}
*/

func main() {
	log.Printf("Connecting to mongodb")

	session, err := mgo.Mongo("localhost")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB("command")

	repo := newRepository()
	repo.rebuild(db)
	defer repo.snapshot(db)

	handler := newCommandHandler(repo, db.C("events"))
	tr := newTaskRest(db.C("tasks"), handler)
	rest.Resource("tasks", tr)

	log.Printf("About to listen on 4041")
	go func() {
		err = http.ListenAndServe(":4041", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	select {
            case sig := <-signal.Incoming:
                fmt.Println("***Caught", sig)
                switch sig.(os.UnixSignal) {
                    case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
                        log.Println("Shutting down...")
                        return
                }
        }
}
