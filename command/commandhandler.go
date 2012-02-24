package main

import (
	"log"
	"bestbuys"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
	"bitbucket.org/Swoogan/mongorest"
)

//type data map[string]interface{}
type handler func(data bestbuys.Data, repo repository) *bestbuys.Event
type handlerPool map[string]handler

type commandHandler struct {
	pool handlerPool
	repo repository
	col  mgo.Collection
}

func newCommandHandler(repo repository, col mgo.Collection) commandHandler {
	pool := handlerPool{
		"createGame":    createGame,
		"setWallet":     setWallet,
		"setUpkeep":     setUpkeep,
		"setBalance":    setBalance,
		"setIncome":     setIncome,
		"setLandIncome": setLandIncome,
	}
	return commandHandler{pool, repo, col}
}

func (c commandHandler) Created(doc mongorest.Document) {
	name := doc["name"].(string)
	if handler, ok := c.pool[name]; ok {
		data := doc["data"].(map[string]interface{})
		edata := bestbuys.Data(data)
		log.Println("Handling command:", name)
		event := handler(edata, c.repo)
		c.store(event)
		dispatch(event)
	} else {
		log.Printf("No handler specified for command: %v", name)
	}
}

func (c commandHandler) store(e *bestbuys.Event) {
	e.Date = bson.Now()

	if err := c.col.Insert(e); err != nil {
		log.Println("Could not save to datastore:", err)
	}
}

//
// HANDLERS
//
func createGame(data bestbuys.Data, repo repository) *bestbuys.Event {
	log.Println("In here")
	id := bson.NewObjectId()
	data["id"] = id.Hex()

	var lands []land
	for _, landData := range data["lands"].([]interface{}) {
		var land land
		for key, value := range landData.(map[string]interface{}) {
			switch key {
			case "name":
				land.Name = value.(string)
			case "cost":
				land.Cost = bestbuys.Money(value.(float64))
			case "income":
				land.Income = bestbuys.Money(value.(float64))
			}
		}
		lands = append(lands, land)
	}

	repo[id.Hex()] = game{
		Id:      id,
		Finance: finance{0, 0},
		Monies:  monies{0, 0, 0},
		Lands:   lands,
	}

	log.Println("Created game:", data["name"])

	return createEvent("gameCreated", data)
}

func setIncome(data bestbuys.Data, repo repository) *bestbuys.Event {
	id, game := getGame(data, repo)
	game.Finance.Income = bestbuys.Money(data["income"].(float64))
	repo[id] = game
	hourly := game.Finance.hourly()
	data["hourly"] = hourly
	data["daily"] = game.Finance.daily(hourly)
	return createEvent("incomeSet", data)
}

func setUpkeep(data bestbuys.Data, repo repository) *bestbuys.Event {
	id, game := getGame(data, repo)
	game.Finance.Upkeep = bestbuys.Money(data["upkeep"].(float64))
	repo[id] = game
	hourly := game.Finance.hourly()
	data["hourly"] = hourly
	data["daily"] = game.Finance.daily(hourly)
	return createEvent("upkeepSet", data)
}

func setBalance(data bestbuys.Data, repo repository) *bestbuys.Event {
	id, game := getGame(data, repo)
	game.Monies.Balance = bestbuys.Money(data["balance"].(float64))
	repo[id] = game
	data["totalMonies"] = game.Monies.total()
	return createEvent("balanceSet", data)
}

func setWallet(data bestbuys.Data, repo repository) *bestbuys.Event {
	id, game := getGame(data, repo)
	game.Monies.Wallet = bestbuys.Money(data["wallet"].(float64))
	repo[id] = game
	data["totalMonies"] = game.Monies.total()
	return createEvent("walletSet", data)
}

func setLandIncome(data bestbuys.Data, repo repository) *bestbuys.Event {
	id, game := getGame(data, repo)
	game.Monies.Lands = bestbuys.Money(data["landIncome"].(float64))
	repo[id] = game
	data["totalMonies"] = game.Monies.total()
	return createEvent("landIncomeSet", data)
}

//
// Helpers
//

func getGame(data bestbuys.Data, repo repository) (string, game) {
	id := data["game"].(string)
	game := repo[id]
	return id, game
}

func createEvent(name string, data bestbuys.Data) *bestbuys.Event {
	return &bestbuys.Event{name, bson.Now(), data}
}
