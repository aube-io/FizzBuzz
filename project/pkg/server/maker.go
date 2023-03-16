package server

import (
	"fmt"
	"os"

	"github.com/anotherhope/FizzBuzz/project/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

var server *Server

func Create() *Server {
	if server == nil {
		config := fiber.Config{
			AppName: config.HOSTNAME,
			Prefork: true,
		}

		if os.Getppid() <= 1 {
			fmt.Println("WARNING: fiber in downgrade mode please use docker run --pid=host")
			config.Prefork = false
		}

		server = &Server{
			app: fiber.New(config),
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
