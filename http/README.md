# HTTP server and client

In these exercises you'll learn to write a small
RESTful server and a client that consumes the API.

## 1 Basic HTTP server

Write a small program that starts an HTTP server
and listens for requests.
When a request is received return a static text.

Send requests to your server to check that it works.

> Tip!
>
> Look at the `net/http` package.


## 2 Basic client

Write a program that connects to your server and
prints the response to the screen.

> Tip!
>
> Put your client in a different folder than your server
> and use `package main` for both of them.

## 3 Return some JSON

Extend your server to return JSON instead of just text.
Create a `Person` struct with the attributes `Name` and
`Age` and have it serialize as JSON.

> Tip!
>
> You are going to need the `encoding/json` package.

**Extra**: Can you also make your server return different
representations based on the `Accept` header?

## 4 Parse the JSON

Make your client able to parse the incoming JSON and represent
it as a value in your program.
Use the data to output the information to the screen.

## 5 POST (and PUT?) data

Add support for `POST`ing data
(for example the `Person` you created earlier).
Store the data in memory for later retrieval.

The server should return the correct HTTP status code
and a generated ID.

Add a handler that can return the correct value based on
an ID in the URL so that the client can read the value again.

**Extra**: Validate the incoming data and return a fitting
HTTP status when the data is invalid
(missing name, negagive age and so on).

**Even more extra**: Can you add support for updating a resource
as well?

## 6 POST (and PUT?) from the client

Implement support for `POST`ing data to your server.
Retreive the data again using the ID the server returned to you.

**Extra**: Have your client program accept command line parameters
that are passed on to the server. Return a non-zero exit code if
the request fails.

## 7 Testing

Write tests for your HTTP handler function using the
`net/http/httptest` package.

Here's something to help you get started:

1. Create a file in the `server` directory and name it `server_test.go`
2. Use this template:
   ```go
   package main

   import (
       "testing"
       "net/http"
       "net/http/httptest"
   )

   func TestGetPerson(t *testing.T) {
       // write your test here.
   }
   ```
3. Create a `http.Request` to represent the type of request you want to
   test (a GET request for a person with a given ID).
4. Use `httptest.NewRecorder()` to create a `ResponseRecorder`
5. Create an HTTP handler that wraps your handler with `http.HandlerFunc()`
   and call it with `.ServeHTTP()` passing your `ResponseRecorder` and request.
6. Inspect the `ResponseRecorder` and test for its properties such as status code
   and body.

Refer to the [the documentation](https://golang.org/pkg/net/http/httptest/).
