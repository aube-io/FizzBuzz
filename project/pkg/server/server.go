// Server provides a simple HTTP/HTTPS server implementation for serving static and connect API.
package server

import (
	"os"

	"github.com/anotherhope/FizzBuzz/project/config"
	"github.com/anotherhope/FizzBuzz/project/internal/api"
	"github.com/anotherhope/FizzBuzz/project/internal/lib"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app      *fiber.App
	api      fiber.Router
	versions map[string]fiber.Router
}

func (server *Server) Start() {
	go server.http()
	if config.EnableHTTPS {
		go server.https()
	}

	os.Exit(lib.WaitStatus())
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
	lib.WithCriticalError(server.app.Listen(":80"))
}

func (server *Server) https() {
	lib.WithCriticalError(server.app.ListenTLS(":443", config.TLSPath+config.TLSCertFile, config.TLSPath+config.TLSKeyFile))
}
