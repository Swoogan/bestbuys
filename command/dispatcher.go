package main

import (
	"log"
	"rpc"
	"eventbus/event"
)

func rpcCall(address string, method string, e *event.Event) {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Println("OMG connect failed need to queue this!1!!")
		return
	}
	defer client.Close()

	var reply int
	err = client.Call(method, e, &reply)
	if err != nil {
		log.Println("OMG call failed, need to queue this!!!")
	}
}

func denormalize(e *event.Event) {
	rpcCall(":4042", "Denormalizer.HandleEvent", e)
}

func schedule(e *event.Event) {
	rpcCall(":4043", "Scheduler.HandleEvent", e)
}

func dispatch(e *event.Event) {
	denormalize(e)
	//	schedule(e)
	log.Printf("Dispatched event: %v", e.Name)
}
