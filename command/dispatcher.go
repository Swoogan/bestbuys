package main

import (
	"log"
)

type Data map[string]interface{}

type Command struct {
	name string
	data Data
}

type handlerPool map[string]func (cmd Command)

type dispatcher struct {
	pool handlerPool
}

func newDispatcher() dispatcher {
	return dispatcher{ make(handlerPool)  }
}

func(d dispatcher) Dispatch(cmd Command) {
	log.Println("Handle command")
}

type Dispatches interface {
	Dispatch(cmd Command)
}
