package main

import (
	"os"
	"log"
	"fmt"
	"http"
	"syscall"
	"os/signal"
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

	mongorest.New(db, "games")

	log.Printf("About to listen on 4040")
	go func() {
		err = http.ListenAndServe(":4040", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	select {
            case sig := <-signal.Incoming:
                fmt.Printf("***Caught %s\n", sig)
                switch sig.(os.UnixSignal) {
                    case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
                        log.Println("Shutting down...")
                        return
                }
        }
}
