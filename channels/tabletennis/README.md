# Table tennis simulator

Simulate a game of table tennis using
goroutines and channels.
Send a ball between to players (goroutines)
via a table (channel).

Experiment with different ways to block the
main function while waiting for the game to complete.

You can use

* pause the function with `time.Sleep`
* use a second channel to signal to the main function
  that the game is complete
* use a [WaitGroup](https://golang.org/pkg/sync/#WaitGroup)
  to wait for a number of hits before continuing.
  