package denormalizer

import (
	"os"
	"log"
)

type Data map[string]interface{}

type Event struct {
	Name string
	Data Data
}

type Denormalizer struct {}

func (d *Denormalizer) HandleEvent(e *Event, i *int) os.Error {
	log.Printf("Handling Event: %v", e.Name)
	return nil
}
