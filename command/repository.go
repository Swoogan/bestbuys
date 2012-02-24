package main

import (
	"os"
	"log"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
)

type repository struct {
	games map[string]game
	log *log.Logger
}

func newRepository(logger *log.Logger) repository {
	return repository{make(map[string]game, 3), logger}
}

func (r repository) rebuild(d mgo.Database) {
	r.log.Println("Loading snapshots")
	var result []game
	err := d.C("snapshots").Find(nil).All(&result)
	if err != nil {
		r.log.Println("Could not load snapshots:", err)
		return
	}

	for _, game := range result {
		r.games[game.Id.Hex()] = game
	}
}

func (r repository) snapshot(d mgo.Database) os.Error {
	r.log.Println("Saving snapshots")
	for id, game := range r.games {
		game.LastSaved = bson.Now() // what if save fails???
		selector := bson.M{"_id": bson.ObjectIdHex(id)}
		_, err := d.C("snapshots").Upsert(selector, bson.M{"$set": game})
		if err != nil {
			r.log.Println("Could not save snapshot:", err)
			return err
		}
	}
	return nil
}
