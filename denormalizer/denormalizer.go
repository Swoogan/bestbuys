package denormalizer

import (
	"os"
	"log"
	"eventbus/event"
)

type Denormalizer struct {}

func (d *Denormalizer) HandleEvent(e *event.Event, i *int) os.Error {
	log.Printf("Handling Event: %v", e.Name)
	return nil
}
