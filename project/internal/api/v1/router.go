package v1

import (
	"github.com/anotherhope/FizzBuzz/project/internal/api"
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
		Namespace: "fizzbuz",
	}
)
