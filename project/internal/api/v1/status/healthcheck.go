package status

import "github.com/gofiber/fiber/v2"

// @Summary		Show the status of server.
// @Description	get the status of server.
// @Tags		Status
// @Accept		*/*
// @Produce		json
// @Success		204	{object}	nil
// @Router		/api/v1/status/healthcheck [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNoContent).SendString("")
}
