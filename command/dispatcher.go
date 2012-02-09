package main

import (
	"log"
	"rpc"
	"denormalizer"
)

type handler func(d denormalizer.Data)
type handlerPool map[string]handler

type command struct {
	name string
	data denormalizer.Data
}

type dispatcher struct {
	pool handlerPool
}

func newDispatcher() dispatcher {
	pool := handlerPool {
			"setWallet": setWallet,
		}
	return dispatcher{ pool }
}

func(d dispatcher) Dispatch(cmd command) {
	if handler, ok := d.pool[cmd.name]; ok {
		handler(cmd.data)
	} else {
		log.Printf("No handler specified for command: %v", cmd.name)
	}
}

type Dispatches interface {
	Dispatch(cmd command)
}

//
// HANDLERS
//

func setWallet(d denormalizer.Data) {
	client, err := rpc.DialHTTP("tcp", ":4042")
	if err != nil {
		log.Println("OMG connect failed need to queue this!1!!:", err)
		return
	}

	event := &denormalizer.Event{"walletWasSet", d}
	var reply int

	err = client.Call("Denormalizer.HandleEvent", event, &reply)
	if err != nil {
		log.Println("OMG call failed, need to queue this!!!")
	}

	log.Printf("Handle Command: %v", "setWallet")
}
