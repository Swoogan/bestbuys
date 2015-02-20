package main

import (
	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/Swoogan/bestbuys/domain"
)

func setIncome(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	logger.Println(game)
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
	game.Finance.Balance = domain.Money(data["balance"].(float64))
	repo[id] = game
	data["totalFinance"] = game.Finance.TotalMoney()
	return createEvent("balanceSet", data)
}

func setWallet(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	game.Finance.Wallet = domain.Money(data["wallet"].(float64))
	repo[id] = game
	data["totalFinance"] = game.Finance.TotalMoney()
	return createEvent("walletSet", data)
}

func setLandIncome(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	game.Finance.Lands = domain.Money(data["landIncome"].(float64))
	repo[id] = game
	data["totalFinance"] = game.Finance.TotalMoney()
	return createEvent("landIncomeSet", data)
}

func setStructureCost(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	name := data["structureName"].(string)
	st := game.Structures[name]
	st.Cost = domain.Money(data["structureCost"].(float64))
	if len(game.Structures) == 0 {
		logger.Println("Structures weren't hydrated correctly.")
		//return nil, error
	}
	game.Structures[name] = st
	repo[id] = game
	return createEvent("structureCostSet", data)
}

func generatePurchases(data domain.Data, repo repository) *domain.Event {
	_, game := getGame(data, repo)
	root := domain.NewTree(len(game.Structures))
	root.Build(game.Structures, game.Finance, numberOfBuys)
	var buys []domain.Buy
	best := root.FindBestPath(numberOfBuys, buys, 0, 0)
	data["purchases"] = best
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
