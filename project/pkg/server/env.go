package server

import (
	"os"

	"github.com/anotherhope/FizzBuzz/project/internal/exit"
)

const (
	certFile = "/.build/certs/server.crt"
	keyFile  = "/.build/certs/server.key"
)

func pwd() string {
	path, err := os.Getwd()

	exit.WithCriticalError(err)

	return path
}
