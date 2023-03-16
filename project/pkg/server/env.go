package server

import (
	"os"

	"github.com/anotherhope/FizzBuzz/project/internal/exit"
)

const (
	certFile = "/certs/server.crt"
	keyFile  = "/certs/server.key"
)

func pwd() string {
	path, err := os.Getwd()

	exit.WithCriticalError(err)

	return path
}
