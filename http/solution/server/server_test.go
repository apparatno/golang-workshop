package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetPerson(t *testing.T) {
	data := make(map[string]Person)
	birthdate, err := time.Parse(time.RFC3339, "1998-11-20T15:38:00Z")
	if err != nil {
		t.Fatal(err)
	}
	data["1"] = Person{Name: "Test Person", Age: 20, Birthdate: birthdate}
	h := makeHandler(data)

	req, err := http.NewRequest("GET", "/person/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status, got %d want %d", status, http.StatusOK)
	}

	expected := `{"id":"1","name":"Test Person","age":20,"birthdate":"1998-11-20T15:38:00Z"}`
	body := strings.TrimSpace(rr.Body.String())
	if body != expected {
		t.Errorf("wrong body, got  %s want %s", body, expected)
	}
}
