package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL, err)
	writeJSONError(w, http.StatusInternalServerError, "the server encounterd a problem")
}

func (app *application) statusBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s error: %s", r.Method, r.URL, err)
	writeJSONError(w, http.StatusInternalServerError, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found error error: %s path: %s error: %s", r.Method, r.URL, err)
	writeJSONError(w, http.StatusNotFound, "not found")
}
