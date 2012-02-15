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
		"createGame": createGame,
		"setWallet":  setWallet,
		"setUpkeep":  setUpkeep,
		"setBalance": setBalance,
		"setIncome":  setIncome,
		"setLandIncome":  setLandIncome,
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
func createGame(data event.Data, repo repository) *event.Event {
	id := bson.NewObjectId()
	data["id"] = id.Hex()
	repo[id.Hex()] = game{ id, finance{0,0}, monies{0,0,0} }
	return &event.Event{"gameCreated", bson.Now(), data}
}

func setIncome(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.Finance.Income = int64(data["income"].(float64))
	repo[id] = game
	hourly := game.Finance.hourly()
	data["hourly"] = hourly
	data["daily"] = game.Finance.daily(hourly)
	return &event.Event{"incomeSet", bson.Now(), data}
}

func setUpkeep(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.Finance.Upkeep = int64(data["upkeep"].(float64))
	repo[id] = game
	hourly := game.Finance.hourly()
	data["hourly"] = hourly
	data["daily"] = game.Finance.daily(hourly)
	return &event.Event{"upkeepSet", bson.Now(), data}
}

func setBalance(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.Monies.Balance = int64(data["balance"].(float64))
	repo[id] = game
	data["totalMonies"] = game.Monies.total()
	return &event.Event{"balanceSet", bson.Now(), data}
}

func setWallet(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.Monies.Wallet = int64(data["wallet"].(float64))
	repo[id] = game
	data["totalMonies"] = game.Monies.total()
	return &event.Event{"walletSet", bson.Now(), data}
}

func setLandIncome(data event.Data, repo repository) *event.Event {
	id := data["game"].(string)
	game := repo[id]
	game.Monies.Lands = int64(data["landIncome"].(float64))
	repo[id] = game
	data["totalMonies"] = game.Monies.total()
	return &event.Event{"landIncomeSet", bson.Now(), data}
}
