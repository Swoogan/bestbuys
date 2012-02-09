package main

import "denormalizer"

func main() {
	denorm := new(denormalizer.Denormalizer)
	rpc.Register(denorm)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
