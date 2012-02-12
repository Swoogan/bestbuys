package denormalizer

import (
	"os"
	"log"
	"eventbus/event"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
)

type handler func(mgo.Database, event.Data) (err os.Error)
type handlerPool map[string]handler

type Denormalizer struct {
	database mgo.Database
	pool     handlerPool
}

func New(database mgo.Database) *Denormalizer {
	pool := handlerPool{
		"walletSet":  walletSet,
		"upkeepSet":  upkeepSet,
		"balanceSet": balanceSet,
		"incomeSet":  incomeSet,
	}

	return &Denormalizer{database, pool}
}

func (d *Denormalizer) HandleEvent(e *event.Event, i *int) os.Error {
	if handler, ok := d.pool[e.Name]; ok {
		return handler(d.database, e.Data)
	}

	log.Printf("No handler specified for event: %v", e.Name)
	return nil
}

// 
// Handlers
//

func walletSet(database mgo.Database, data event.Data) (err os.Error) {
	log.Println("Handling Event: walletSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
			"wallet": data["wallet"], 
			"total": data["total"],
			}}
	if err = database.C("games").Update(selector, change); err != nil {
		log.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func upkeepSet(database mgo.Database, data event.Data) os.Error {
	log.Println("Handling Event: upkeepSet")
	return nil
}

func balanceSet(database mgo.Database, data event.Data) os.Error {
	log.Println("Handling Event: balanceSet")
	return nil
}

func incomeSet(database mgo.Database, data event.Data) os.Error {
	log.Println("Handling Event: incomeSet")
	return nil
}
