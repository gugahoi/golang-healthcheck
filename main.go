package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// VERSION is the application version
var VERSION = "sample-version"

// CommitSHA is the latet commit that will be populated at build time
var CommitSHA = "123-123-123"

// Main simply starts and http server with the /healthchek endpoint
func main() {
	port := ":80"
	log.Printf("Listening on %v", port)
	if err := http.ListenAndServe(port, handler()); err != nil {
		panic(err)
	}
}

// Handler associates the routes and their respective functions
func handler() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/healthcheck", healthFunc)
	return h
}

// Healthfunc simple writes a '200 OK' header
func healthFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(Information{
		Description: "Web API for ANZ",
		Commit:      CommitSHA,
		Version:     VERSION,
	})
	if err != nil {
		// How to test this scenario?
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(body)
}

// Information is the body structure of the healthcheck endpoint
type Information struct {
	Version     string `json:"version"`
	Commit      string `json:"commit"`
	Description string `json:"description"`
}
