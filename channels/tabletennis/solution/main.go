package main

import (
	"fmt"
	"sync"
	"time"
)

type ball struct {
	hits int
}

func main() {
	b := ball{hits: 0}
	table := make(chan ball)

	// A waitgroup is good for signaling when a
	// goroutine is done
	w := sync.WaitGroup{}
	w.Add(10) // we want ten hits

	go player(table, &w)
	go player(table, &w)

	table <- b

	w.Wait() // block this routine until the waitgroup reaches zero
}

func player(table chan ball, w *sync.WaitGroup) {
	// wait for the ball
	for b := range table {
		time.Sleep(300 * time.Millisecond)
		if b.hits == 10 {
			continue
		}

		b.hits++ // hit the ball
		fmt.Printf("hit %d\n", b.hits)
		w.Done()   // decrements the waitgroup
		table <- b // pass the ball back on the table
	}
}
