package main

import (
	"log"
	"rpc"
	"eventbus/event"
)

type handler func(d event.Data)
type handlerPool map[string]handler

type command struct {
	name string
	data event.Data
}

type dispatcher struct {
	pool handlerPool
}

func newDispatcher() dispatcher {
	pool := handlerPool {
			"setWallet": setWallet,
			"setUpkeep": setUpkeep,
			"setBalance": setBalance,
			"setIncome": setIncome,
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

func rpcCall(address string, method string, e *event.Event) {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Println("OMG connect failed need to queue this!1!!:", err)
		return
	}

	var reply int
	err = client.Call(method, e, &reply)
	if err != nil {
		log.Println("OMG call failed, need to queue this!!!")
	}
	client.Close()
}

func denormalize(e *event.Event) {
	rpcCall(":4042", "Denormalizer.HandleEvent", e)
}

func schedule(e *event.Event) {
	rpcCall(":4043", "Scheduler.HandleEvent", e)
}

func dispatch(e *event.Event) {
	denormalize(e)
	schedule(e)
	log.Printf("Dispatched event: %v", e.Name)
}

//
// HANDLERS
//

func setWallet(d event.Data) {
	event := &event.Event{"walletWasSet", d}
	dispatch(event)
}


func setUpkeep(d event.Data) {
	event := &event.Event{"upkeepWasSet", d}
	dispatch(event)
}

func setBalance(d event.Data) {
	event := &event.Event{"balanceWasSet", d}
	dispatch(event)
}

func setIncome(d event.Data) {
	event := &event.Event{"incomeWasSet", d}
	dispatch(event)
}
