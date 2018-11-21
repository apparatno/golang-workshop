# Getting started

Some quick tips to get started:

* Keep each solution in a separate folder
* Call your file `main.go`.
  It's not strictly necessary, just good practice.
* Always use `package main` and create a `main()` function
* Try to compile your programs with `go build`.
  Name your program with the `-o` flag
* ... but you don't have to compile it if you only use one file.
  Run it directly with `go run main.go`.

### Testing

Familiarize yourself with testing in Go by writing tests for your
solutions to to the exercises
(except for Hello World, perhaps...).

Typically to test functions in a file named `foo.go`
you create a file named `foo_test.go`.
Give it the following structure:

```go
package main

import "testing"

func TestMyThing(t *testing.T) {
    // implement your test here
}
```

Run your test with `go test .`.


## Hello world

The foundation of all programming:
write a program that prints "Hello, World!"
to the screen.

## Sum

Given a list of numbers
return the sum of all elements in the list.

> Tip!
>
> Check out `range`...

## esreveR

Reverse any input string and print it to the screen.

## Word count

Implement a program that can take a text and count the words in it.

> Tip!
>
> A `map` could be useful here.

**Extra**

* Sort the result by most occurences first.
* Read the text from a file on disk or via HTTP.

## Structs & methods

Create a representation of geometric shapes
and implement `area()` and `perimeter()`
as methods on them.

**Extra** Define an interface that describes these behaviors
and implement a function that can accept the interface and
print information about the shape to the screen.

## Caesar cipher (ROT13)

Implement the Caesar cipher in a way so that you can both
encrypt and decrypt a string.

The Caesar cipher works like this:
Given a key `n`, which is a number between 1 and 25,
encrypt a text by moving each letter `n` places forward
in the alphabet.
Decrypting is done by moving `n` places backward.
