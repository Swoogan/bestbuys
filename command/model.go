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
	var result []game
	err := d.C("snapshot").Find(nil).All(&result)
	if err != nil {
		log.Println("Could not load snapshots:", err)
		return
	}

	for _, game := range result {
		r[game.id.Hex()] = game
	}
}

func (r repository) snapshot(d mgo.Database) os.Error {
	log.Println("in snapshot")
	for id, game := range r {
		selector := bson.M{"_id": id}
		_, err := d.C("snapshot").Upsert(selector, game)
		if err != nil {
			log.Println("Could not save snapshot:", err)
			return err
		}
	}
	return nil
}

type game struct {
	id bson.ObjectId `bson: "_id"`
	finance finance
	monies  monies
}

type finance struct {
	income int64
	upkeep int64
}

func (f finance) hourly() int64 {
	return f.income - f.upkeep
}

func (f finance) daily(hourly int64) int64 {
	return hourly * 24
}

type monies struct {
	balance int64
	wallet  int64
	lands   int64
}

func (m monies) total() int64 {
	return m.balance + m.wallet + m.lands
}
