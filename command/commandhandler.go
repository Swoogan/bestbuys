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
type handler func(data event.Data, repo repository) *event.Event
type handlerPool map[string]handler

type commandHandler struct {
	pool handlerPool
	repo repository
}

func newCommandHandler(repo repository) commandHandler {
	pool := handlerPool{
		"setWallet":  setWallet,
		"setUpkeep":  setUpkeep,
		"setBalance": setBalance,
		"setIncome":  setIncome,
	}
	return commandHandler{pool,repo}
}

func (c commandHandler) Handle(cmd command) {
	if handler, ok := c.pool[cmd.name]; ok {
		data := event.Data {
			"name": cmd.data["name"], 
			"data": cmd.data["data"],
		}
		event := handler(data, c.repo)
		//store(event)
		dispatch(event)
	} else {
		log.Printf("No handler specified for command: %v", cmd.name)
	}
}

type HandlesCommand interface {
	Handle(cmd command)
}

//
// HANDLERS
//
func setIncome(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.finance.income = data["income"].(int64)
	repo[id] = game
	hourly := game.finance.hourly()
	data["hourly"] = hourly
	data["daily"] = game.finance.daily(hourly)
	return &event.Event{"incomeSet", data}
}

func setUpkeep(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.finance.upkeep = data["upkeep"].(int64)
	repo[id] = game
	hourly := game.finance.hourly()
	data["hourly"] = hourly
	data["daily"] = game.finance.daily(hourly)
	return &event.Event{"upkeepSet", data}
}

func setBalance(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.monies.balance = data["balance"].(int64)
	repo[id] = game
	data["total"] = game.monies.total()
	return &event.Event{"balanceSet", data}
}

func setWallet(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.monies.wallet = data["wallet"].(int64)
	repo[id] = game
	data["total"] = game.monies.total()
	return &event.Event{"walletSet", data}
}

func setLand(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.monies.lands = data["lands"].(int64)
	repo[id] = game
	data["total"] = game.monies.total()
	return &event.Event{"landSet", data}
}
