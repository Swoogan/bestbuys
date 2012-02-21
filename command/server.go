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

func main() {
	flag.Parse()

	logger := createLogger(*logfile)

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
	tasks := mongorest.Resource{DB: db, Name: "tasks", Handler: handler}
	mongorest.Attach(tasks, logger)

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

func createLogger(logfile string) *log.Logger {
        output := os.Stderr
        if logfile != "" {
                var err os.Error
                output, err = os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
                if err != nil {
                        log.Fatal(err)
                }
        }
        return log.New(output, "", log.LstdFlags)
}

