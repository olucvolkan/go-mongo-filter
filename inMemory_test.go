package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInMemoryHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/in-memory?key=abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(inMemoryGetHandler())
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "abc"
	var inMemoryResponseHandler InMemoryResponse
	json.Unmarshal([]byte(rr.Body.String()), &inMemoryResponseHandler)
	if inMemoryResponseHandler.Key != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
