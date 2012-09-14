package main

import (
	"domain"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"bitbucket.org/Swoogan/mongorest"
)

const numberOfBuys = 2

type handler func(domain.Data, repository) *domain.Event
type handlerPool map[string]handler

type commandHandler struct {
	pool handlerPool
	repo repository
	col  mgo.Collection
}

func newCommandHandler(repo repository, col mgo.Collection) commandHandler {
	pool := handlerPool{
		"createGame":        createGame,
		"setWallet":         setWallet,
		"setUpkeep":         setUpkeep,
		"setBalance":        setBalance,
		"setIncome":         setIncome,
		"setLandIncome":     setLandIncome,
		"setStructureCost":  setStructureCost,
		"generatePurchases": generatePurchases,
	}
	return commandHandler{pool, repo, col}
}

func (c commandHandler) Created(doc mongorest.Document) {
	name := doc["name"].(string)
	if handler, ok := c.pool[name]; ok {
		data := doc["data"].(map[string]interface{})
		edata := domain.Data(data)
		logger.Println("Handling command:", name)
		event := handler(edata, c.repo)
		c.store(event)
		dispatch(event)
	} else {
		logger.Println("No handler specified for command:", name)
	}
}

func (c commandHandler) store(e *domain.Event) {
	e.Date = bson.Now()

	if err := c.col.Insert(e); err != nil {
		logger.Println("Could not save to datastore:", err)
	}
}
