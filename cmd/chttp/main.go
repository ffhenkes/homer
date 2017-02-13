package main

import (
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/NeowayLabs/logger"
	"github.com/ffhenkes/homer/feedback"
	"github.com/ffhenkes/homer/routes"
	"github.com/julienschmidt/httprouter"
)

func main() {

	var logH = logger.Namespace("homer.http")

	var addr = os.Getenv("HOMER_ADDRESS")
	var mongoAddr = os.Getenv("HOMER_MONGODB_ADDRESS")

	// Instantiate a new router
	r := httprouter.New()

	nr := routes.NewRouter(
		getSession(mongoAddr),
	)

	defer nr.CloseAll()

	// health check
	r.GET("/v1", nr.Health)

	// odyssey
	r.POST("/odysseys", nr.CreateOdyssey)

	r.GET("/odysseys", nr.GetAllOdyssey)

	r.GET("/odysseys/:id", nr.GetOdyssey)

	r.PATCH("/odysseys/:id", nr.PatchOdyssey)

	r.DELETE("/odysseys/:id", nr.RemoveOdyssey)

	r.DELETE("/odysseys", nr.RemoveAllOdyssey)

	// step
	r.POST("/steps", nr.CreateStep)

	r.GET("/steps", nr.GetAllStep)

	r.GET("/steps/:id", nr.GetStep)

	r.PATCH("/steps/:id", nr.PatchStep)

	r.DELETE("/steps/:id", nr.RemoveStep)

	r.DELETE("/steps", nr.RemoveAllStep)

	// quest
	r.POST("/quests", nr.CreateQuest)

	r.GET("/quests", nr.GetAllQuests)

	r.GET("/quests/:id", nr.GetQuest)

	r.DELETE("/quests/:id", nr.RemoveQuest)

	r.DELETE("/quests", nr.RemoveAllQuests)

	logH.Info("Running on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		logH.Fatal("ListenAndServe: %s", err)
	}

}

func getSession(mongoAddr string) *mgo.Session {
	session, err := mgo.Dial(mongoAddr)
	feedback.FailOnError(err, "Failed to connect to MongoDB!")
	return session
}
