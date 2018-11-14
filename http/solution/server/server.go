package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Person struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Birthdate time.Time `json:"birthdate"`
}

func main() {
	// A map to store our data in
	data := make(map[string]Person)

	// Initialize storage with a dummy
	data["1"] = Person{
		ID:        "1",
		Name:      "Ricco",
		Age:       39,
		Birthdate: time.Date(1979, 3, 2, 21, 25, 0, 0, time.UTC),
	}

	// Create the HTTP handler and start listening
	handler := makeHandler(data)
	http.HandleFunc("/person/", handler)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

// makeHandler is a closure around the actual handler function.
// It makes the data repository available to the handler function.
func makeHandler(data map[string]Person) func(http.ResponseWriter, *http.Request) {
	// The returned function satisfies the http.HandleFunc interface
	return func(w http.ResponseWriter, r *http.Request) {
		var result Person

		// Detect the HTTP method
		switch r.Method {
		case "GET":
			// Get data based on the ID in the path
			p, err := get(data, r)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
				return
			}
			result = p
		case "POST":
			// Accept the incoming data and save it to the repository
			p, err := save(data, r)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte(err.Error()))
				return
			}
			result = p
		}

		// If we got this far everything went well. Encode the data
		// as JSON
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		e := json.NewEncoder(w)
		if err := e.Encode(result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

// get handles GET requests. It accepts the data repository and the request.
// It returns a Person or an error if the ID wasn't found.
func get(data map[string]Person, r *http.Request) (Person, error) {
	ID := idFromPath(r)
	p, ok := data[ID]
	if !ok {
		return Person{}, fmt.Errorf("ID %s not found", ID)
	}
	p.ID = ID
	return p, nil
}

// idFromPath crudely parses the path and finds the ID.
func idFromPath(r *http.Request) string {
	path := r.RequestURI
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return ""
	}
	candidate := parts[len(parts)-1]
	if candidate != "person" { // ¯\_(ツ)_/¯
		return candidate
	}
	return ""
}

// save handles POST requests. It accepts the data repository and the request.
// It decodes the request body as a Person and stores it to the repository
// or returns an error if the body cannot be decoded.
func save(data map[string]Person, r *http.Request) (Person, error) {
	defer r.Body.Close() // Always remember to close readers! `defer` is gold...
	ID := generateID(data)
	result := Person{} // Write the decoded JSON into this value.
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&result); err != nil {
		return Person{}, fmt.Errorf("invalid person: %v", err)
	}

	if err := validate(result); err != nil {
		return Person{}, err
	}

	result.ID = ID
	data[ID] = result
	return result, nil
}

// generateID generates a new ID based on the length of the repository.
func generateID(data map[string]Person) string {
	items := len(data)
	items++
	return fmt.Sprintf("%d", items)
}

func validate(p Person) error {
	if p.Name == "" {
		return fmt.Errorf("name cannot be blank")
	}
	if p.Age < 0 {
		return fmt.Errorf("age cannot be negative")
	}
	if p.Birthdate.After(time.Now()) {
		return fmt.Errorf("birthdate cannot be in the future")
	}
	return nil
}
