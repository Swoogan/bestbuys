package main

import (
	"denormalizer"
	"rpc"
	"net"
	"log"
	"http"
)

func main() {
	denorm := new(denormalizer.Denormalizer)
	rpc.Register(denorm)
	rpc.HandleHTTP()
	log.Println("About to listen on localhost:4042")
	l, e := net.Listen("tcp", ":4042")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
