package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Creates a new quest
func (router *Router) CreateQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Get all quests
func (router *Router) GetAllQuests(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Get a quest by id
func (router *Router) GetQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Updates partially a quest
func (router *Router) PatchQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Removes all quests
func (router *Router) RemoveAllQuests(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Removes one quest
func (router *Router) RemoveQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}
