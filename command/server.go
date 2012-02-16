package main

import (
	"os"
	"log"
	"fmt"
	"flag"
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
	mongo := flag.String("m", "localhost", "Mongo server address")
	dbname := flag.String("d", "command", "Mongo database name")
	address := flag.String("a", ":4041", "Address to listen on")
	flag.Parse()

	log.Println("Connecting to mongodb")

	session, err := mgo.Mongo(*mongo)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB(*dbname)

	repo := newRepository()
	repo.rebuild(db)
	defer repo.snapshot(db)

	handler := newCommandHandler(repo, db.C("events"))
	tr := newTaskRest(db.C("tasks"), handler)
	rest.Resource("tasks", tr)

	log.Printf("About to listen on %v", *address)
	go func() {
		err = http.ListenAndServe(*address, nil)
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
