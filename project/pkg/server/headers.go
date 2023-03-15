package server

import (
	"github.com/gofiber/fiber/v2"
)

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
