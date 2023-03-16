package server

import (
	"os"

	"github.com/anotherhope/FizzBuzz/project/config"
	"github.com/anotherhope/FizzBuzz/project/internal/api"
	"github.com/anotherhope/FizzBuzz/project/internal/exit"
	"github.com/gofiber/fiber/v2"
)

const (
	certFile = "/server.crt"
	keyFile  = "/server.key"
)

type Server struct {
	app      *fiber.App
	api      fiber.Router
	versions map[string]fiber.Router
}

func (server *Server) Start() {
	go server.http()
	go server.https()

	os.Exit(exit.WaitStatus())
}

func (server *Server) API(api *api.API) {
	api.Register(server.version(api).Group(api.Namespace))
}

func (server *Server) version(api *api.API) fiber.Router {
	if router, exist := server.versions[api.Version]; exist {
		return router
	}

	server.versions[api.Version] = server.api.Group(api.Version)

	return server.versions[api.Version]
}

func (server *Server) http() {
	exit.WithCriticalError(server.app.Listen(":80"))
}

func (server *Server) https() {
	exit.WithCriticalError(server.app.ListenTLS(":443", config.TLS_PATH+config.TLS_CERT_FILE, config.TLS_PATH+config.TLS_KEY_FILE))
}
