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

func WithCriticalError(criticalError error) {
	if criticalError != nil {
		errors <- criticalError // J'aurais pu utiliser un panic(criticalError) mais je trouve ca moins élégants que de sortir de cette manière
	}
}

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
