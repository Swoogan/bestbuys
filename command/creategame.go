package main

import (
        "gopkg.in/mgo.v2/bson"

        "bitbucket.org/Swoogan/bestbuys/domain"
)

func createGame(data domain.Data, repo repository) *domain.Event {
	id := bson.NewObjectId()
	data["id"] = id.Hex()

	lands := make(map[string]*domain.Land)
	for _, landData := range data["lands"].([]interface{}) {
		var land domain.Land
		for key, value := range landData.(map[string]interface{}) {
			switch key {
			case "name":
				land.Name = value.(string)
			case "cost":
				land.Cost = domain.Money(value.(float64))
			case "income":
				land.Income = domain.Money(value.(float64))
			case "retainAlways":
				land.RetainAlways = value.(bool)
			}
		}
		lands[land.Name] = &land
	}

	structures := make(domain.Structures)
	for _, sData := range data["structures"].([]interface{}) {
		var structure domain.Structure
		for key, value := range sData.(map[string]interface{}) {
			switch key {
			case "name":
				structure.Name = value.(string)
			case "cost":
				structure.Cost = domain.Money(value.(float64))
			case "increase":
				structure.Increase = domain.Money(value.(float64))
			case "income":
				structure.Income = domain.Money(value.(float64))
			case "builtOn":
				land := value.(string)
				structure.BuiltOn = lands[land]
			}
		}
		structures[structure.Name] = structure
	}

	repo[id.Hex()] = domain.Game{
		Id:         id,
		Finance:    domain.Finance{0, 0, 0, 0, 0},
		Structures: structures,
	}

	logger.Println("Created game:", data["name"])

	return createEvent("gameCreated", data)
}
