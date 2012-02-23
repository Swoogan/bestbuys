package denormalizer

import (
	"os"
	"log"
	"bestbuys"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
)

type handler func(mgo.Database, bestbuys.Data) (err os.Error)
type handlerPool map[string]handler

type Denormalizer struct {
	database mgo.Database
	pool     handlerPool
}

func New(database mgo.Database) *Denormalizer {
	pool := handlerPool{
		"gameCreated":   gameCreated,
		"walletSet":     walletSet,
		"upkeepSet":     upkeepSet,
		"balanceSet":    balanceSet,
		"incomeSet":     incomeSet,
		"landIncomeSet": landIncomeSet,
	}

	return &Denormalizer{database, pool}
}

func (d *Denormalizer) HandleEvent(e *bestbuys.Event, i *int) os.Error {
	if handler, ok := d.pool[e.Name]; ok {
		return handler(d.database, e.Data)
	}

	log.Printf("No handler specified for event: %v", e.Name)
	return nil
}

// 
// Handlers
//

func gameCreated(database mgo.Database, data bestbuys.Data) (err os.Error) {
	log.Println("Handling Event: gameCreated")
	log.Println("Game id is:", data["id"])
	id := bson.ObjectIdHex(data["id"].(string))
	info := bson.M{
		"_id":         id,
		"name":        data["name"],
		"lands": data["lands"],
		"wallet":      0,
		"balance":     0,
		"landIncome":  0,
		"totalMonies": 0,
		"income":      0,
		"upkeep":      0,
		"hourly":      0,
		"daily":       0,
	}
	if err = database.C("games").Insert(info); err != nil {
		log.Println("Could not insert game in the datastore, ", err, ": ", data["name"])
	}
	return

}

func incomeSet(database mgo.Database, data bestbuys.Data) (err os.Error) {
	log.Println("Handling Event: incomeSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"income": data["income"],
		"hourly": data["hourly"],
		"daily":  data["daily"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		log.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func upkeepSet(database mgo.Database, data bestbuys.Data) (err os.Error) {
	log.Println("Handling Event: upkeepSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"upkeep": data["upkeep"],
		"hourly": data["hourly"],
		"daily":  data["daily"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		log.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func walletSet(database mgo.Database, data bestbuys.Data) (err os.Error) {
	log.Println("Handling Event: walletSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"wallet":      data["wallet"],
		"totalMonies": data["totalMonies"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		log.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func balanceSet(database mgo.Database, data bestbuys.Data) (err os.Error) {
	log.Println("Handling Event: balanceSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"balance":     data["balance"],
		"totalMonies": data["totalMonies"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		log.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func landIncomeSet(database mgo.Database, data bestbuys.Data) (err os.Error) {
	log.Println("Handling Event: landIncomeSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"landIncome":  data["landIncome"],
		"totalMonies": data["totalMonies"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		log.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}
