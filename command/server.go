package main

import (
	"os"
	"fmt"
	"log"
	"flag"
	"http"
	"syscall"
	"bestbuys"
	"os/signal"
	"launchpad.net/mgo"
	"bitbucket.org/Swoogan/mongorest"
)

/*
func main() {
	StoreTest()
	LoadTest()
}
*/
var mongo *string = flag.String("m", "localhost", "Mongo server address")
var dbname *string = flag.String("d", "command", "Mongo database name")
var address *string = flag.String("a", ":4041", "Address to listen on")
var logfile *string = flag.String("l", "", "Log file to write to")

var logger *log.Logger

func main() {
	flag.Parse()

	logger = bestbuys.NewLogger(*logfile, "Command Handler\t")

	logger.Println("Connecting to mongodb")
	session, err := mgo.Mongo(*mongo)
	if err != nil {
		logger.Fatal(err)
		return
	}
	defer session.Close()

	logger.Printf("Opening database %v", *dbname)
	db := session.DB(*dbname)

	repo := newRepository()
	repo.rebuild(db)
	defer repo.snapshot(db)

	handler := newCommandHandler(repo, db.C("events"))
	commands := mongorest.Resource{DB: db, Name: "commands", Handler: handler}
	mongorest.Attach(commands, logger)

	logger.Printf("About to listen on %v", *address)
	go func() {
		err = http.ListenAndServe(*address, nil)
		if err != nil {
			logger.Fatal(err)
		}
	}()

	select {
	case sig := <-signal.Incoming:
		fmt.Println("***Caught", sig)
		switch sig.(os.UnixSignal) {
		case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
			logger.Println("Shutting down...")
			return
		}
	}
}
