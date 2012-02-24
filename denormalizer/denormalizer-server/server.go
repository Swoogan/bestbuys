package main

import (
	"os"
	"rpc"
	"net"
	"flag"
	"http"
	"syscall"
	"bestbuys"
	"os/signal"
	"denormalizer"
	"launchpad.net/mgo"
)

var mongo *string = flag.String("m", "localhost", "Mongo server address")
var dbname *string = flag.String("d", "query", "Mongo database name")
var address *string = flag.String("a", ":4042", "Address to listen on")
var logfile *string = flag.String("l", "", "File to log to")

func main() {
	flag.Parse()

	logger := bestbuys.NewLogger(*logfile)

	logger.Printf("Connecting to mongodb")
	session, err := mgo.Mongo(*mongo)
	if err != nil {
		logger.Fatal(err)
		return
	}
	defer session.Close()

	db := session.DB(*dbname)

	denorm := denormalizer.New(db)

	rpc.Register(denorm)
	rpc.HandleHTTP()
	logger.Printf("About to listen on %v", *address)

	go func() {
		l, err := net.Listen("tcp", *address)
		if err != nil {
			logger.Fatal("listen error:", err)
		}
		http.Serve(l, nil)
	}()

	select {
	case sig := <-signal.Incoming:
		logger.Println("***Caught", sig)
		switch sig.(os.UnixSignal) {
		case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
			logger.Println("Shutting down...")
			return
		}
	}
}
