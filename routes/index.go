package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

type (
	Router struct {
		session *mgo.Session
	}
)

func NewRouter(s *mgo.Session) *Router {
	return &Router{s}
}

// Health This one is just a simple sanity check
func (router *Router) Health(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "“Of all creatures that breathe and move upon the earth, nothing is bred that is weaker than man.” - Homer, The Odyssey\n")
}

func (router *Router) CloseAll() {
	router.session.Close()
}
