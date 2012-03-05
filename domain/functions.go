package domain

import (
	"os"
	"log"
)

func NewLogger(logfile string, prefix string) *log.Logger {
	output := os.Stderr
	if logfile != "" {
		var err os.Error
		output, err = os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
	return log.New(output, prefix, log.LstdFlags)
}
