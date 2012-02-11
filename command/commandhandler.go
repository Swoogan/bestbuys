package main

import (
	"log"
	"eventbus/event"
)

type command struct {
	name string
	data data
}

type data map[string]interface{}
type handler func(d event.Data) *event.Event
type handlerPool map[string]handler

type commandHandler struct {
	pool handlerPool
}

func newCommandHandler() commandHandler {
	pool := handlerPool{
		"setWallet":  setWallet,
		"setUpkeep":  setUpkeep,
		"setBalance": setBalance,
		"setIncome":  setIncome,
	}
	return commandHandler{pool}
}

func (c commandHandler) Handle(cmd command) {
	if handler, ok := c.pool[cmd.name]; ok {
		event := handler(event.Data {
			"name": cmd.data["name"], 
			"data": cmd.data["data"],
		})
		dispatch(event)
	} else {
		log.Printf("No handler specified for command: %v", cmd.name)
	}
}

type HandlesCommand interface {
	Handle(cmd command)
}

var games []game

//
// HANDLERS
//
func setIncome(d event.Data) *event.Event {
	//hourly := d["income"] - 
	return &event.Event{"incomeSet", d}
}

func setUpkeep(d event.Data) *event.Event {
	//hourly := d.
	return &event.Event{"upkeepSet", d}
}

func setBalance(d event.Data) *event.Event {
	return &event.Event{"balanceSet", d}
}

func setWallet(d event.Data) *event.Event {
	return &event.Event{"walletSet", d}
}

func setLand(d event.Data) *event.Event {
	return &event.Event{"landSet", d}
}
