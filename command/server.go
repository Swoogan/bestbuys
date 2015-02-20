package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"bitbucket.org/Swoogan/mongorest"
	"gopkg.in/mgo.v2"

	"bitbucket.org/Swoogan/bestbuys/domain"
)

var mongo *string = flag.String("m", "localhost", "Mongo server address")
var dbname *string = flag.String("d", "command", "Mongo database name")
var address *string = flag.String("a", ":4041", "Address to listen on")
var logfile *string = flag.String("l", "", "Log file to write to")

var logger *log.Logger

func main() {
	flag.Parse()

	logger = domain.NewLogger(*logfile, "Command Handler\t")

	logger.Println("Connecting to mongodb")
	session, err := mgo.Dial(*mongo)
	if err != nil {
		logger.Fatal(err)
		return
	}
	defer session.Close()

	logger.Printf("Opening database '%v'", *dbname)
	db := session.DB(*dbname)

	repo := newRepository()
	repo.rebuild(db)

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

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	for sig := range c {
		repo.snapshot(db)
		logger.Printf("Received %v, shutting down...", sig)
		os.Exit(1)
	}
}
