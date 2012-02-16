package main

import (
	"os"
	"fmt"
	"rpc"
	"net"
	"log"
	"http"
	"syscall"
	"os/signal"
	"denormalizer"
	"launchpad.net/mgo"
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

	denorm := denormalizer.New(db)

	rpc.Register(denorm)
	rpc.HandleHTTP()
	log.Println("About to listen on localhost:4042")

	go func() {
		l, err := net.Listen("tcp", ":4042")
		if err != nil {
			log.Fatal("listen error:", err)
		}
		http.Serve(l, nil)
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
