package denormalizer

import (
	"os"
	"log"
	"labix.org/v2/mgo"
	"bestbuys_go/domain"
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
		"gameCreated":        gameCreated,
		"walletSet":          walletSet,
		"upkeepSet":          upkeepSet,
		"balanceSet":         balanceSet,
		"incomeSet":          incomeSet,
		"landIncomeSet":      landIncomeSet,
		"structureCostSet":   structureCostSet,
		"purchasesGenerated": purchasesGenerated,
	}

	return &EventHandler{database, pool, logger}
}

func (d *EventHandler) HandleEvent(e *domain.Event, i *int) os.Error {
	if handler, ok := d.pool[e.Name]; ok {
		return handler(d.database, e.Data, d.log)
	}

	d.log.Println("No handler specified for event:", e.Name)
	return nil
}
