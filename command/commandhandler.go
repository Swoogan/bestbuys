package main

/*
import (
	"fmt"
)
*/

type Handler struct {
}

/*
func NewHandler() *Handler {
	return &Handler{ }
}
*/

func(h Handler) handleCommand(cmd Command) {
}

type commandHandler interface {
	handleCommand(cmd Command)
}
