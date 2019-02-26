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
	log.Printf("Listening on %v", port)
}

// Main simply starts and http server with the /healthchek endpoint
func main() {
	if err := http.ListenAndServe(port, handler()); err != nil {
		log.Fatal(err)
	}
}

// Handler associates the routes and their respective functions
func handler() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/healthcheck", healthFunc)
	return h
}

// Healthfunc '200 OK' and replies with some basic information regarding the application
func healthFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body, err := json.Marshal(ApplicationInfo{
		Name: []HealthcheckBody{
			{
				Description: "Web API for ANZ",
				Commit:      CommitSHA,
				Version:     Version,
			},
		},
	})
	if err != nil {
		// TODO: How to test this scenario?
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(body)
}

// ApplicationInfo is the top level application information
type ApplicationInfo struct {
	Name []HealthcheckBody `json:"golang-healthchecker"`
}

// HealthcheckBody is the body structure of the healthcheck endpoint
type HealthcheckBody struct {
	Version     string `json:"version"`
	Commit      string `json:"lastcommitsha"`
	Description string `json:"description"`
}
