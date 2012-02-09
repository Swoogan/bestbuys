package main

import (
	"denormalizer"
	"rpc"
	"net"
	"log"
	"http"
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

	l, err := net.Listen("tcp", ":4042")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	http.Serve(l, nil)
}
