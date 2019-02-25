package main

import "net/http"

// Main simply starts and http server with the /healthchek endpoint
func main() {
	if err := http.ListenAndServe(":80", handler()); err != nil {
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
}
