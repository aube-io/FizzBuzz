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

		app := fiber.New(config)

		server = &Server{
			app:      app,
			api:      app.Group("api"),
			versions: make(map[string]fiber.Router),
		}

		server.app.Use(setSecurityHeaders)
		server.app.Get("/docs/*", swagger.HandlerDefault)
		server.app.Group("/api", setRedirectOnEntryPointAPI)
	}

	return server
}

func setRedirectOnEntryPointAPI(c *fiber.Ctx) error {
	if c.Path() == "/api" || c.Path() == "/api/" {
		return c.Redirect("/docs")
	}
	return c.Next()
}

func setSecurityHeaders(c *fiber.Ctx) error {
	// Activer HSTS (HTTP Strict Transport Security)
	c.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	// Activer CSP (Content Security Policy)
	c.Set("Content-Security-Policy", "default-src 'unsafe-inline' 'self' fonts.gstatic.com fonts.googleapis.com;img-src data: 'self'")
	// Activer CORS (Cross-Origin Resource Sharing)
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET,POST,HEAD,PUT,DELETE,PATCH")
	c.Set("Access-Control-Allow-Headers", "*")
	c.Set("Access-Control-Allow-Credentials", "true")

	return c.Next()
}
