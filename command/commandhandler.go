package main

import (
	"domain"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
	"bitbucket.org/Swoogan/mongorest"
)

type handler func(domain.Data, repository) *domain.Event
type handlerPool map[string]handler

type commandHandler struct {
	pool handlerPool
	repo repository
	col  mgo.Collection
}

func newCommandHandler(repo repository, col mgo.Collection) commandHandler {
	pool := handlerPool{
		"createGame":       createGame,
		"setWallet":        setWallet,
		"setUpkeep":        setUpkeep,
		"setBalance":       setBalance,
		"setIncome":        setIncome,
		"setLandIncome":    setLandIncome,
		"setStructureCost": setStructureCost,
		"generatePurchases": generatePurchases,
	}
	return commandHandler{pool, repo, col}
}

func (c commandHandler) Created(doc mongorest.Document) {
	name := doc["name"].(string)
	if handler, ok := c.pool[name]; ok {
		data := doc["data"].(map[string]interface{})
		edata := domain.Data(data)
		logger.Println("Handling command:", name)
		event := handler(edata, c.repo)
		c.store(event)
		dispatch(event)
	} else {
		logger.Println("No handler specified for command:", name)
	}
}

func (c commandHandler) store(e *domain.Event) {
	e.Date = bson.Now()

	if err := c.col.Insert(e); err != nil {
		logger.Println("Could not save to datastore:", err)
	}
}

//
// HANDLERS
//
func generatePurchases(data domain.Data, repo repository) *domain.Event {
	id, game := getGame(data, repo)
	// do stuff here
	repo[id] = game
	return createEvent("purchasesGenerated", data)
}

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
			}
		}
		lands[land.Name] = &land
	}

	var structures []domain.Structure
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
		structures = append(structures, structure)
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
	game.Monies.Lands = domain.Money(data["structureCost"].(float64))
	repo[id] = game
	return createEvent("structureCostSet", data)
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
