package denormalizer

import (
	"os"
	"log"
	"domain"
	"launchpad.net/mgo"
)

type handler func(mgo.Database, domain.Data, *log.Logger) (err os.Error)
type handlerPool map[string]handler

type EventHandler struct {
	database mgo.Database
	pool     handlerPool
	log      *log.Logger
}

func New(database mgo.Database, logger *log.Logger) *EventHandler {
	pool := handlerPool{
		"gameCreated":      gameCreated,
		"walletSet":        walletSet,
		"upkeepSet":        upkeepSet,
		"balanceSet":       balanceSet,
		"incomeSet":        incomeSet,
		"landIncomeSet":    landIncomeSet,
		"structureCostSet": structureCostSet,
		"purchasesGenerated": purchasesGenerated,
	}

	return &EventHandler{database, pool, logger}
}

func (eh *EventHandler) HandleEvent(e *domain.Event, i *int) os.Error {
	if handler, ok := eh.pool[e.Name]; ok {
		return handler(eh.database, e.Data, eh.log)
	}

	eh.log.Println("No handler specified for event:", e.Name)
	return nil
}
