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
	return &Denormalizer{
		database,
		handlerPool{
			"walletWasSet":  walletWasSet,
			"upkeepWasSet":  upkeepWasSet,
			"balanceWasSet": balanceWasSet,
			"incomeWasSet":  incomeWasSet,
		},
	}
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

func walletWasSet(database mgo.Database, data event.Data) (err os.Error) {
	log.Println("Handling Event: walletWasSet")
	selector := bson.M{"_id": bson.ObjectIdHex("4f18697b3946caf4198dd697")}
	change := bson.M{"$set": bson.M{"wallet": data["wallet"]}}
	if err = database.C("games").Update(selector, change); err != nil {
		log.Println("Could not update the datastore")
		log.Println(err)
	}
	return
}

func upkeepWasSet(database mgo.Database, data event.Data) os.Error {
	log.Println("Handling Event: upkeepWasSet")
	return nil
}

func balanceWasSet(database mgo.Database, data event.Data) os.Error {
	log.Println("Handling Event: balanceWasSet")
	return nil
}

func incomeWasSet(database mgo.Database, data event.Data) os.Error {
	log.Println("Handling Event: incomeWasSet")
	return nil
}
