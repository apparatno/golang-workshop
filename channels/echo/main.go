package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	// Create a channel to pass text
	c := make(chan string)

	// Create a channel to receive signals on
	s := make(chan os.Signal)

	// Create a wait group...
	var wg sync.WaitGroup
	// ...and initialize it to 1
	wg.Add(1)

	// Handle signals from the OS
	signal.Notify(s, os.Interrupt, os.Kill)
	go func() {
		_ = <-s
		close(c)
		// Decrease the initial work to wait for
		// to allow the program to exit
		wg.Done()
	}()

	// Start some workers
	for i := 0; i < 5; i++ {
		go echo(c, &wg)
	}

	// Show the prompt and wait for input
	go prompt(c, &wg)

	// Block here until all work is complete
	wg.Wait()
}

// prompt displays a prompt and waits for input.
// When input is received it is passed on to the
// channel and the function waits for more input.
func prompt(c chan string, wg *sync.WaitGroup) {
	r := bufio.NewReader(os.Stdin)
	for {
		if c == nil {
			// The channel is closed so there's nothing
			// more to do here
			return
		}
		fmt.Print("Enter text: ")
		s, _ := r.ReadString('\n')
		// Tell the wait group that we are working on
		// +1 items of work
		wg.Add(1)
		// Pass the work to a worker via the channel
		c <- s
	}
}

// echo listens to a channel and "works" before
// displaying the received text on the screen.
func echo(c chan string, wg *sync.WaitGroup) {
	// iterate over the channel
	for s := range c {
		// pretend to work
		time.Sleep(5 * time.Second)
		// print
		fmt.Println(s)
		// decrease the work
		wg.Done()
	}
}
