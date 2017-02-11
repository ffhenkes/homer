package feedback

import (
	"log"

	"github.com/NeowayLabs/logger"
)

var logF = logger.Namespace("homer.feedback.error")

// generic fatal error treatment
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatal("%s: %s", msg, err)
	}
}
