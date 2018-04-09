package main

import (
	"time"
	"flag"

	"github.com/getamis/sirius/log"
)



func main() {
	delay := flag.Int("delay", 1000, "millisecond")
	flag.Parse()

	ticker := time.NewTicker(time.Duration(*delay) *time.Millisecond)
	logger := log.New()
	for {
		for t := range ticker.C {
			logger.Error("Long string testing, the message and keypairs with one space. SHOW!", "tickat", t, "hash", "0x12345", "previous", "no", "after", "true", "drop", "event:sandbox", "events", "bank did added")
			logger.Trace("Short string, with more spaces", "time", t)
			logger.Info("Two words in the key", "tick at", t)
			logger.Warn("Allowed all characters !@#$%^&*()-?=\\/", "tick at", t)
			logger.Debug("Key values showed in message, key=values")
			logger.Info("Invalid key value showed in message, key= value")
			logger.Info("value not in key, key= ")
			logger.Info("value not in key, append new kv, key= ", "key1", "value2")
		}
	}
}