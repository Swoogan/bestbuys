package main

import (
	"os"
	"log"
	"net"
	"flag"
	"net/rpc"
	"os/signal"
	"net/rpc/jsonrpc"

        "gopkg.in/mgo.v2"

        "bitbucket.org/Swoogan/bestbuys/domain"
        "bitbucket.org/Swoogan/bestbuys/denormalizer"
)

var mongo *string = flag.String("m", "localhost", "Mongo server address")
var dbname *string = flag.String("d", "query", "Mongo database name")
var address *string = flag.String("a", ":4042", "Address to listen on")
var logfile *string = flag.String("l", "", "File to log to")

var logger *log.Logger

func main() {
	flag.Parse()

	logger = domain.NewLogger(*logfile, "Denormalizer\t")

	logger.Printf("Connecting to mongodb")
	session, err := mgo.Dial(*mongo)
	if err != nil {
		logger.Fatal(err)
	}
	defer session.Close()

	db := session.DB(*dbname)
	eh := denormalizer.New(db, logger)

	rpc.Register(eh)
	go serve()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for sig := range c {
		logger.Printf("Received %v, shutting down...", sig)
		os.Exit(1)
	}
}

func serve() {
	logger.Printf("About to listen on %v", *address)
	l, err := net.Listen("tcp", *address)
	if err != nil {
		logger.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Println(err)
			return
		}

		go jsonrpc.ServeConn(conn)
	}
}
