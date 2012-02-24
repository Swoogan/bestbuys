package main

import (
	"net"
	"bestbuys"
	"rpc/jsonrpc"
)

func rpcCall(address string, method string, e *bestbuys.Event) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		logger.Println("OMG connect failed need to queue this!1!!")
		return
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)
	defer client.Close()

	var reply int
	if err = client.Call(method, e, &reply); err != nil {
		logger.Println("OMG call failed, need to queue this!!!")
	}
}

func denormalize(e *bestbuys.Event) {
	rpcCall(":4042", "Denormalizer.HandleEvent", e)
}

func schedule(e *bestbuys.Event) {
	rpcCall(":4043", "Scheduler.HandleEvent", e)
}

func dispatch(e *bestbuys.Event) {
	denormalize(e)
	//	schedule(e)
	logger.Printf("Dispatched event: %v", e.Name)
}
