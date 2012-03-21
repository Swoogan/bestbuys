package main

import (
	"os"
	"log"
	"rpc"
	"net"
	"flag"
	"syscall"
	"domain"
	"os/signal"
	"rpc/jsonrpc"
	"denormalizer"
	"launchpad.net/mgo"
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
	session, err := mgo.Mongo(*mongo)
	if err != nil {
		logger.Fatal(err)
	}
	defer session.Close()

	db := session.DB(*dbname)
	eh := denormalizer.New(db, logger)

	rpc.Register(eh)
	go serve()

	for {
		select {
		case sig := <-signal.Incoming:
			logger.Println("***Caught", sig)
			switch sig.(os.UnixSignal) {
			case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
				logger.Println("Shutting down...")
				return
			default:
				continue
			}
		}
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
