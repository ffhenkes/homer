package routes

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Creates a new odyssey
func (router *Router) CreateOdyssey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Get all  odysseys
func (router *Router) GetAllOdyssey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Get an odyssey by id
func (router *Router) GetOdyssey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Updates partially an odyssey
func (router *Router) PatchOdyssey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Removes all odysseys
func (router *Router) RemoveAllOdyssey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}

// Removes one odyssey
func (router *Router) RemoveOdyssey(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO only a stub

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", "stub")
}
