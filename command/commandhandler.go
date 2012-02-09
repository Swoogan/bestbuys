package main

/*
import (
	"fmt"
)
*/

type handlerPool map[string]func (cmd Command)

type Disbatcher struct {
	pool handlerPool
}

/*
func NewHandler() *Handler {
	return &Handler{ }
}
*/

func(d Disbatcher) handleCommand(cmd Command) {
}

type commandHandler interface {
	handleCommand(cmd Command)
}
