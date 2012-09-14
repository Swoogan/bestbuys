package main

import (
	"bestbuys_go/domain"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type repository map[string]domain.Game

func newRepository() repository {
	return make(map[string]domain.Game, 3)
}

func (r repository) rebuild(d mgo.Database) {
	logger.Println("Loading snapshots")
	var result []domain.Game
	err := d.C("snapshots").Find(nil).All(&result)
	if err != nil {
		logger.Println("Could not load snapshots:", err)
		return
	}

	for _, game := range result {
		r[game.Id.Hex()] = game
	}
}

func (r repository) snapshot(d mgo.Database) error {
	logger.Println("Saving snapshots")
	for id, game := range r {
		game.LastSaved = bson.Now() // what if save fails???
		selector := bson.M{"_id": bson.ObjectIdHex(id)}
		_, err := d.C("snapshots").Upsert(selector, bson.M{"$set": game})
		if err != nil {
			logger.Println("Could not save snapshot:", err)
			return err
		}
	}
	return nil
}
