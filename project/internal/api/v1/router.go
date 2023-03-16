// API Grouping flagged by version number 1
package v1

import (
	"github.com/anotherhope/FizzBuzz/project/internal/api"
	"github.com/anotherhope/FizzBuzz/project/internal/api/v1/fizzbuzz"
	"github.com/anotherhope/FizzBuzz/project/internal/api/v1/status"
	"github.com/gofiber/fiber/v2"
)

var (
	STATUS = &api.API{
		Version:   "v1",
		Namespace: "status",
		Get: map[string][]func(*fiber.Ctx) error{
			"healthcheck": {status.HealthCheck},
		},
	}

	FIZZBUZZ = &api.API{
		Version:   "v1",
		Namespace: "fizzbuzz",
		Get: map[string][]func(*fiber.Ctx) error{
			"/:int1/:int2/:limit/:str1/:str2": {fizzbuzz.FizzBuzzControls, fizzbuzz.FizzBuzzHits, fizzbuzz.FizzBuzz},
			"/stats":                          {fizzbuzz.Stats},
		},
	}
)
