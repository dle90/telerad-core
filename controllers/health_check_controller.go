package controllers

import (
	"telerad-core-module/internals/responses"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "Telerad Core Success", time.Now().String()))
}
