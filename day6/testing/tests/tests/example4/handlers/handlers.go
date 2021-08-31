// Package handlers provides the endpoints for the web service.
package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes sets the routes for the web service.
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON returns a simple JSON document.
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@goinggo.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

// NotCoverSendJSON returns a simple JSON document.
func NotCoverSendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@goinggo.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
