package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Version is the application version
var Version = "sample-version"

// CommitSHA is the latet commit that will be populated at build time
var CommitSHA = "123-123-123"
var port string

func init() {
	// Port could be read from ENV variable or from flags later on
	port = ":80"
}

// Main simply starts and http server with the /healthchek endpoint
func main() {
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
	body, err := json.Marshal(HealthcheckBody{
		Description: "Web API for ANZ",
		Commit:      CommitSHA,
		Version:     Version,
	})
	if err != nil {
		// TODO: How to test this scenario?
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(body)
}

// HealthcheckBody is the body structure of the healthcheck endpoint
type HealthcheckBody struct {
	Version     string `json:"version"`
	Commit      string `json:"lastcommitsha"`
	Description string `json:"description"`
}
