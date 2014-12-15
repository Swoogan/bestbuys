package main

import (
	"bitbucket.org/Swoogan/mongorest"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/Swoogan/bestbuys/domain"
)

const numberOfBuys = 2

type handler func(domain.Data, repository) *domain.Event
type handlerPool map[string]handler

type commandHandler struct {
	pool handlerPool
	repo repository
	col  *mgo.Collection
}

func newCommandHandler(repo repository, col *mgo.Collection) commandHandler {
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
		data, ok := doc["data"].(map[string]interface{})
		if ok {
			edata := domain.Data(data)
			logger.Println("Handling command:", name)
			event := handler(edata, c.repo)
			c.store(event)
			dispatch(event)
		} else {
			logger.Println("Received data is not a map[string]interface{}: ", data)
		}
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
