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
		selector := bson.M{"_id": bson.ObjectIdHex(id)}
		_, err := d.C("snapshots").Upsert(selector, bson.M{"$set": game})
		if err != nil {
			log.Println("Could not save snapshot:", err)
			return err
		}
	}
	return nil
}

type game struct {
	Id        bson.ObjectId "hack" // this should be "_id" but then Upsert doesn't do anything
	Finance   finance
	Monies    monies
	LastSaved bson.Timestamp
}

type finance struct {
	Income int64
	Upkeep int64
}

func (f finance) hourly() int64 {
	return f.Income - f.Upkeep
}

func (f finance) daily(hourly int64) int64 {
	return hourly * 24
}

type monies struct {
	Balance int64
	Wallet  int64
	Lands   int64
}

func (m monies) total() int64 {
	return m.Balance + m.Wallet + m.Lands
}
