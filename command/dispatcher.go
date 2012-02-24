package main

import (
	"log"
	"rpc"
	"bestbuys"
)

func rpcCall(address string, method string, e *bestbuys.Event) {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Println("OMG connect failed need to queue this!1!!")
		return
	}
	defer client.Close()

	var reply int
	if err = client.Call(method, e, &reply); err != nil {
		log.Println("OMG call failed, need to queue this!!!")
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
	log.Printf("Dispatched event: %v", e.Name)
}
