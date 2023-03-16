package server

import (
	"os"

	"github.com/anotherhope/FizzBuzz/project/internal/exit"
)

const (
	certFile = "/etc/ssl/certs/server.crt"
	keyFile  = "/etc/ssl/certs/server.key"
)

func pwd() string {
	path, err := os.Getwd()

	exit.WithCriticalError(err)

	return path
}
