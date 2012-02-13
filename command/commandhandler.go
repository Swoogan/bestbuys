package main

import (
	"log"
	"eventbus/event"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
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
	col mgo.Collection
}

func newCommandHandler(repo repository, col mgo.Collection) commandHandler {
	pool := handlerPool{
		"setWallet":  setWallet,
		"setUpkeep":  setUpkeep,
		"setBalance": setBalance,
		"setIncome":  setIncome,
	}
	return commandHandler{pool, repo, col}
}

func (c commandHandler) Handle(cmd command) {
	if handler, ok := c.pool[cmd.name]; ok {
		data := event.Data(cmd.data)
		event := handler(data, c.repo)
		c.store(event)
		dispatch(event)
	} else {
		log.Printf("No handler specified for command: %v", cmd.name)
	}
}

func (c commandHandler) store(e *event.Event) {
	e.Date = bson.Now()

	if err := c.col.Insert(e); err != nil {
		log.Println("Could not save to datastore:", err)
	}
}

type HandlesCommand interface {
	Handle(cmd command)
}

//
// HANDLERS
//
func setIncome(data event.Data, repo repository) *event.Event {
	id, game := getGame(data, repo)
	game.finance.income = int64(data["income"].(float64))
	repo[id] = game
	hourly := game.finance.hourly()
	data["hourly"] = hourly
	data["daily"] = game.finance.daily(hourly)
	return createEvent("incomeSet", data)
}

func setUpkeep(data event.Data, repo repository) *event.Event {
	id, game := getGame(data, repo)
	game.finance.upkeep = int64(data["upkeep"].(float64))
	repo[id] = game
	hourly := game.finance.hourly()
	data["hourly"] = hourly
	data["daily"] = game.finance.daily(hourly)
	return createEvent("upkeepSet", data)
}

func setBalance(data event.Data, repo repository) *event.Event {
	id, game := getGame(data, repo)
	game.monies.balance = int64(data["balance"].(float64))
	repo[id] = game
	data["total"] = game.monies.total()
	return createEvent("balanceSet", data)
}

func setWallet(data event.Data, repo repository) *event.Event {
	id, game := getGame(data, repo)
	game.monies.wallet = int64(data["wallet"].(float64))
	repo[id] = game
	data["total"] = game.monies.total()
	return createEvent("walletSet", data)
}

func setLand(data event.Data, repo repository) *event.Event {
	id, game := getGame(data, repo)
	game.monies.lands = int64(data["lands"].(float64))
	repo[id] = game
	data["total"] = game.monies.total()
	return createEvent("landSet", data)
}

//
// Helpers
//

func getGame(data event.Data, repo repository) (string, game) {
	id := data["game"].(string)
	game := repo[id]
	return id, game
}

func createEvent(name string, data event.Data) *event.Event {
	return &event.Event{name, bson.Now(), data}
}
