package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Person struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Birthdate time.Time `json:"birthdate"`
}

func main() {
	// Command line flags
	// Usage is:
	// client get <ID>
	// client -name <name> -age <age> -birthdate <birthdate> save
	var name string
	var age int
	var birthdate string

	flag.StringVar(&name, "name", "", "name of the person")
	flag.IntVar(&age, "age", -1, "age of the person")
	flag.StringVar(&birthdate, "birthdate", "", "person's birthdate")

	flag.Parse()

	// Get the non-flag arguments (ie the 'get' or 'save' command)
	rest := flag.Args()

	// Die if no command was given
	if len(rest) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	command := rest[0]

	// Figure out what command to run
	switch command {
	case "get":
		ID := flag.Arg(1)
		if err := get(ID); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	case "save":
		// we need to parse the birthdate into a time.Time value
		// parsing dates is a little funky to newcomers
		// see https://golang.org/pkg/time/#example_Parse
		bdate, err := time.Parse(time.RFC3339, birthdate)
		if err != nil {
			fmt.Printf("Invalid birthdate '%s': %v\n", birthdate, err)
			os.Exit(1)
		}
		p := Person{
			Name:      name,
			Age:       age,
			Birthdate: bdate,
		}
		if err = save(p); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	default:
		fmt.Printf("Invalid command '%s'\n", command)
		flag.Usage()
		os.Exit(1)
	}
}

func get(ID string) error {
	// Create a request, leave the third argument (body) as nil
	request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/person/%s", ID), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	// Send the request and handle any errors
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	// Handle HTTP status-level errors
	if response.StatusCode >= 300 {
		return fmt.Errorf("http error: %d (%s)", response.StatusCode, response.Status)
	}
	// We can read the body now, so remember to close it
	defer response.Body.Close()

	// Empty value to read the body into
	person := Person{}
	// JSON decoder that reads the body
	decoder := json.NewDecoder(response.Body)
	// Decode into the person value
	if err = decoder.Decode(&person); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	// Success
	fmt.Printf("Person found: %s is %d years old and has a birthday at %v\n", person.Name, person.Age, person.Birthdate)
	return nil
}

func save(person Person) error {
	// Encode the person as bytes...
	var body bytes.Buffer
	// ... of JSON
	encoder := json.NewEncoder(&body)
	if err := encoder.Encode(&person); err != nil {
		return fmt.Errorf("error encoding data: %v", err)
	}

	// Createt the request. Use the bytes from before as the body
	request, err := http.NewRequest("POST", "http://localhost:8080/person/", &body)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Send the request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}

	// it's safe to read the response. We need to read it to get the ID or any error messages
	defer response.Body.Close()

	// Handle any HTTP status-level errors
	if response.StatusCode >= 300 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		return fmt.Errorf("http error: %d (%s)", response.StatusCode, buf.String())
	}

	// Read into a person value
	person = Person{}
	// create a JSON decoder that reads the body
	decoder := json.NewDecoder(response.Body)
	// Decode
	if err = decoder.Decode(&person); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	// Success
	fmt.Printf("Person created with ID %s\n", person.ID)
	return nil
}
