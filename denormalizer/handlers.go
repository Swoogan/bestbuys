package denormalizer

import (
	"log"
	"bestbuys_go/domain"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func gameCreated(database *mgo.Database, data domain.Data, logger *log.Logger) (err error) {
	logger.Println("Handling Event: gameCreated")
	logger.Println("Game id is:", data["id"])
	id := bson.ObjectIdHex(data["id"].(string))
	info := bson.M{
		"_id":         id,
		"name":        data["name"],
		"lands":       data["lands"],
		"structures":  data["structures"],
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
		logger.Println("Could not insert game in the datastore, ", err, ": ", data["name"])
	}
	return

}

func incomeSet(database *mgo.Database, data domain.Data, logger *log.Logger) (err error) {
	logger.Println("Handling Event: incomeSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"income": data["income"],
		"hourly": data["hourly"],
		"daily":  data["daily"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		logger.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func upkeepSet(database *mgo.Database, data domain.Data, logger *log.Logger) (err error) {
	logger.Println("Handling Event: upkeepSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"upkeep": data["upkeep"],
		"hourly": data["hourly"],
		"daily":  data["daily"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		logger.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func walletSet(database *mgo.Database, data domain.Data, logger *log.Logger) (err error) {
	logger.Println("Handling Event: walletSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"wallet":      data["wallet"],
		"totalMonies": data["totalMonies"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		logger.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func balanceSet(database *mgo.Database, data domain.Data, logger *log.Logger) (err error) {
	logger.Println("Handling Event: balanceSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"balance":     data["balance"],
		"totalMonies": data["totalMonies"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		logger.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func landIncomeSet(database *mgo.Database, data domain.Data, logger *log.Logger) (err error) {
	logger.Println("Handling Event: landIncomeSet")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"landIncome":  data["landIncome"],
		"totalMonies": data["totalMonies"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		logger.Println("Could not update the datastore, ", err, ": ", data["game"])
	}
	return
}

func structureCostSet(database *mgo.Database, data domain.Data, logger *log.Logger) (err error) {
	logger.Println("Handling Event: structureCostSet")
	id := bson.ObjectIdHex(data["game"].(string))
	name := data["structureName"].(string)
	selector := bson.M{"_id": id, "structures.name": name}
	change := bson.M{"$set": bson.M{
		"structures.$.cost": data["structureCost"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		logger.Println("Could not update the datastore, ", err, ": ", selector)
	}
	return
}

func purchasesGenerated(database *mgo.Database, data domain.Data, logger *log.Logger) (err error) {
	logger.Println("Handling Event: purchasesGenerated")
	id := bson.ObjectIdHex(data["game"].(string))
	selector := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{
		"purchases": data["purchases"],
	}}
	if err = database.C("games").Update(selector, change); err != nil {
		logger.Println("Could not update the datastore, ", err, ": ", selector)
	}
	return nil
}
