package main

import (
	"domain"
	"launchpad.net/gobson/bson"
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
			case "retainalways":
				land.RetainAlways = value.(bool)
			}
		}
		lands[land.Name] = &land
	}

	structures := make(map[string]domain.Structure)
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
		Finance:    domain.Finance{0, 0},
		Monies:     domain.Monies{0, 0, 0},
		Structures: structures,
	}

	logger.Println("Created game:", data["name"])

	return createEvent("gameCreated", data)
}

func setIncome(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	game.Finance.Income = domain.Money(data["income"].(float64))
	repo[id] = game
	hourly := game.Finance.Hourly()
	data["hourly"] = hourly
	data["daily"] = game.Finance.Daily(hourly)
	return createEvent("incomeSet", data)
}

func setUpkeep(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	game.Finance.Upkeep = domain.Money(data["upkeep"].(float64))
	repo[id] = game
	hourly := game.Finance.Hourly()
	data["hourly"] = hourly
	data["daily"] = game.Finance.Daily(hourly)
	return createEvent("upkeepSet", data)
}

func setBalance(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	game.Monies.Balance = domain.Money(data["balance"].(float64))
	repo[id] = game
	data["totalMonies"] = game.Monies.Total()
	return createEvent("balanceSet", data)
}

func setWallet(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	game.Monies.Wallet = domain.Money(data["wallet"].(float64))
	repo[id] = game
	data["totalMonies"] = game.Monies.Total()
	return createEvent("walletSet", data)
}

func setLandIncome(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	game.Monies.Lands = domain.Money(data["landIncome"].(float64))
	repo[id] = game
	data["totalMonies"] = game.Monies.Total()
	return createEvent("landIncomeSet", data)
}

func setStructureCost(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	name := data["structureName"].(string)
	st := game.Structures[name]
	st.Cost = domain.Money(data["structureCost"].(float64))
	game.Structures[name] = st
	repo[id] = game
	return createEvent("structureCostSet", data)
}

func generatePurchases(data domain.Data, repo repository) *domain.Event {
	_, game := getGame(data, repo)
	root := domain.NewRootNode(len(game.Structures), game.Finance, game.Monies)
	domain.CreateNodes(root, game.Structures, numberOfBuys)

	best := domain.FindBestChild(root, numberOfBuys, "", 0, 0)
	logger.Println("best:", best)
	//data["purchases"] = best

	return createEvent("purchasesGenerated", data)
}

//
// Helpers
//

func getGame(data domain.Data, repo repository) (string, domain.Game) {
	id := data["game"].(string)
	game := repo[id]
	return id, game
}

func createEvent(name string, data domain.Data) *domain.Event {
	return &domain.Event{name, bson.Now(), data}
}
