package main

import (
	"os"

	"github.com/NeowayLabs/logger"
)

func main() {

	var logA = logger.Namespace("homer.agent")

	var addr = os.Getenv("HOMER_ADDRESS")
	var nsqAddr = os.Getenv("HOMER_NSQ_ADDRESS")

}
