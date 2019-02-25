package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	expectedStatusCode := http.StatusOK
	srv := httptest.NewServer(handler())
	defer srv.Close()
	rsp, err := http.Get(fmt.Sprintf("%s/healthcheck", srv.URL))
	if err != nil {
		t.Fatalf("could not send get request: %v", err)
	}

	if rsp.StatusCode != expectedStatusCode {
		t.Fatalf("expected status code to be %v, got %v", expectedStatusCode, rsp.StatusCode)
	}
}

func TestHealthFunc(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "localhost:80/healthcheck", nil)
	if err != nil {
		t.Fatalf("unable to create request: %v", err)
	}
	expectedResponseCode := http.StatusOK

	healthFunc(w, r)
	if w.Code != expectedResponseCode {
		t.Fatalf("Expected response code to be '%v', got '%v'", expectedResponseCode, w.Code)
	}
}
