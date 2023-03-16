// Grouping various libraries
package lib

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var sigs = make(chan os.Signal, 1)
var errors = make(chan error, 1)
var done = make(chan int, 1)

// WithCriticalError quit the program following when a critical error occurs
func WithCriticalError(criticalError error) {
	if criticalError != nil {
		errors <- criticalError // J'aurais pu utiliser un panic(criticalError) mais je trouve ca moins élégants que de sortir de cette manière
	}
}

// WaitStatus waits for a signal or an error and returns an integer value.
// It registers the SIGINT and SIGTERM signals using the signal.Notify method,
// then waits for either a signal or an error using a select statement.
// If a signal is received, it sends a value of 0 to the "done" channel.
// If an error is received, it prints the error and sends a value of 1 to the "done" channel.
// Finally, it waits for a value to be received from the "done" channel and returns it as the result of the function.
func WaitStatus() int {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigs:
		done <- 0
	case err := <-errors:
		fmt.Println(err)
		done <- 1
	}

	return <-done
}
