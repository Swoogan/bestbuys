package main

import (
//	"os"
//	"io/ioutil"
	"log"
	"fmt"
	"http"
	"json"
	"github.com/Swoogan/rest.go"
	"launchpad.net/mgo"
	"launchpad.net/gobson/bson"
)

var formatting = "Valid JSON is required\n"

type task struct {
	Id bson.ObjectId `json:",omitempty" bson:"_id"`
	Name string
	Data Data
}

type TaskRest struct {
	col mgo.Collection
	handle func(c Command)
}

/*

// Get all of the documents in the mongo collection 
func (mr *TaskRest) Index(w http.ResponseWriter, r *http.Request) {
	var lookup map[string]interface{}
	if len(r.URL.RawQuery) > 0 {
		var err os.Error
		if lookup, err = parseQuery(r.URL.Query()); err != nil {
			rest.BadRequest(w, err.String())
			return
		}
	}

	var result []task
	err := mr.col.Find(lookup).Limit(100).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	switch accept := r.Header.Get("accept"); {
	case strings.Contains(accept, "application/json"):
		enc := json.NewEncoder(w)
		w.Header().Set("content-type", "application/json")
		enc.Encode(&result)
	default:
		w.WriteHeader(http.StatusNotAcceptable)
	}
}

// Find a document in the collection, identified by the ID
func (mr *TaskRest) Find(w http.ResponseWriter, idString string, r *http.Request) {
	var result map[string]interface{}
	id := createIdLookup(idString)
	if err := mr.col.Find(id).One(&result); err != nil {
		rest.NotFound(w)
		return
	}

	switch accept := r.Header.Get("accept"); {
	case strings.Contains(accept, "application/json"):
		enc := json.NewEncoder(w)
		w.Header().Set("content-type", "application/json")
		enc.Encode(&result)
	default:
		w.WriteHeader(http.StatusNotAcceptable)
	}
}
*/

// Create and add a new document to the collection
func (tr *TaskRest) Create(w http.ResponseWriter, r *http.Request) {
	//TODO: Check the content-type
	dec := json.NewDecoder(r.Body)
	var result task
	if err := dec.Decode(&result); err != nil {
		rest.BadRequest(w, formatting)
	        log.Println("Could not decode json")
	        log.Println(err)
		return
	}

	result.Id = bson.NewObjectId()

	if err := tr.col.Insert(result); err != nil {
		rest.BadRequest(w, "Could not insert document")
	        log.Println("Could not save to datastore")
	        log.Println(err)
		return
	}

	output := fmt.Sprintf("%v%v", r.URL.String(), result.Id.Hex())
	tr.handle(Command { result.Name, result.Data })
	rest.Created(w, output)
}

func NewTaskRest(col mgo.Collection, handler func(c Command) ) *TaskRest {
        return &TaskRest{ col, handler }
}

