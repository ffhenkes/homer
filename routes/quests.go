package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Creates a new odyssey
func (router *Router) CreateQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Get all  odysseys
func (router *Router) GetAllQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Get an odyssey by id
func (router *Router) GetQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Updates partially an odyssey
func (router *Router) PatchQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Removes all odysseys
func (router *Router) RemoveAllQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Removes one odyssey
func (router *Router) RemoveQuest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}
