package denormalizer

import (
	"os"
	"log"
	"domain"
	"launchpad.net/mgo"
)

type handler func(mgo.Database, domain.Data, *log.Logger) (err os.Error)
type handlerPool map[string]handler

type Denormalizer struct {
	database mgo.Database
	pool     handlerPool
	log      *log.Logger
}

func New(database mgo.Database, logger *log.Logger) *Denormalizer {
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

	return &Denormalizer{database, pool, logger}
}

func (d *Denormalizer) HandleEvent(e *domain.Event, i *int) os.Error {
	if handler, ok := d.pool[e.Name]; ok {
		return handler(d.database, e.Data, d.log)
	}

	d.log.Println("No handler specified for event:", e.Name)
	return nil
}
