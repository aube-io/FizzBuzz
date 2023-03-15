package server

import (
	"github.com/anotherhope/FizzBuzz/project/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

var server *Server

func Create() *Server {
	if server == nil {
		server = &Server{
			app: fiber.New(fiber.Config{
				AppName: config.HOSTNAME,
				Prefork: true,
			}),
		}

		server.app.Use(setSecurityHeaders)
		server.app.Get("/docs/*", swagger.HandlerDefault)
		server.app.Group("/api", func(c *fiber.Ctx) error {
			if c.Path() == "/api" || c.Path() == "/api/" {
				return c.Redirect("/docs")
			}
			return c.Next()
		})

	}

	return server
}
