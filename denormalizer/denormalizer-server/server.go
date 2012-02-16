package main

import (
	"os"
	"fmt"
	"rpc"
	"net"
	"log"
	"flag"
	"http"
	"syscall"
	"os/signal"
	"denormalizer"
	"launchpad.net/mgo"
)

func main() {
	mongo := flag.String("m", "localhost", "Mongo server address")
	dbname := flag.String("d", "query", "Mongo database name")
	address := flag.String("a", ":4042", "Address to listen on")
	flag.Parse()

	log.Printf("Connecting to mongodb")

	session, err := mgo.Mongo(*mongo)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB(*dbname)

	denorm := denormalizer.New(db)

	rpc.Register(denorm)
	rpc.HandleHTTP()
	log.Printf("About to listen on %v", *address)

	go func() {
		l, err := net.Listen("tcp", *address)
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
