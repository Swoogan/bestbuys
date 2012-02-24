package main

import (
	"os"
	"log"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
)

type repository map[string]game

func newRepository() repository {
	return make(map[string]game, 3)
}

func (r repository) rebuild(d mgo.Database) {
	log.Println("Loading snapshots")
	var result []game
	err := d.C("snapshots").Find(nil).All(&result)
	if err != nil {
		log.Println("Could not load snapshots:", err)
		return
	}

	for _, game := range result {
		r[game.Id.Hex()] = game
	}
}

func (r repository) snapshot(d mgo.Database) os.Error {
	log.Println("Saving snapshots")
	for id, game := range r {
		game.LastSaved = bson.Now() // what if save fails???
		selector := bson.M{"_id": bson.ObjectIdHex(id)}
		_, err := d.C("snapshots").Upsert(selector, bson.M{"$set": game})
		if err != nil {
			log.Println("Could not save snapshot:", err)
			return err
		}
	}
	return nil
}
