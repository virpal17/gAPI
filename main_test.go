package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/welcome", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(greetHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"Greeting":"Hello, World!"}`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	var msg Message
	err = json.Unmarshal(rr.Body.Bytes(), &msg)
	if err != nil {
		t.Fatal(err)
	}

	if msg.Greeting != "Hello, World!" {
		t.Errorf("handler returned wrong message: got %v want %v", msg.Greeting, "Hello, World!")
	}
}
