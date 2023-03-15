package server

import (
	"os"

	"github.com/anotherhope/FizzBuzz/project/api"
	"github.com/anotherhope/FizzBuzz/project/internal/exit"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func (server *Server) Start() {
	go server.http()
	go server.https()

	os.Exit(exit.WaitStatus())
}

func (server *Server) API(api *api.API) {
	api.Register(server.app.Group(api.Version))
}

func (server *Server) http() {
	exit.WithCriticalError(server.app.Listen(":80"))
}

func (server *Server) https() {
	exit.WithCriticalError(server.app.ListenTLS(":443", pwd()+certFile, pwd()+keyFile))
}
