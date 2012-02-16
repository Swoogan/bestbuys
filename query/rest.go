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

func main() {
	var mongo *string = flag.String("m", "localhost", "Mongodb address")
	var dbname *string = flag.String("d", "query", "Mongodb database name")
	var colname *string = flag.String("c", "games", "Mongodb collection name")
	var address *string = flag.String("a", ":4040", "Address to listen on")
	flag.Parse()

	log.Printf("Connecting to mongodb")

	session, err := mgo.Mongo(*mongo)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	db := session.DB(*dbname)

	mongorest.New(db, *colname)

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
