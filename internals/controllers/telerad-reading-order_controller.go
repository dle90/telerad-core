package controllers

import (
	"net/http"

	"telerad-core-module/internals/responses"
	"telerad-core-module/internals/secure"
	"telerad-core-module/internals/services"

	teleradReadingOrderControllerRequests "telerad-core-module/internals/requests/telerad-reading-order-controller_requests"

	_error "telerad-core-module/error"

	"github.com/BeeTechHub/go-common/logger"
	commonUtils "github.com/BeeTechHub/go-common/utils"

	"github.com/gofiber/fiber/v2"
)

func PartnerCreateReadingOrder(c *fiber.Ctx) error {
	logger.Info("PartnerCreateReadingOrder starting....")

	var request teleradReadingOrderControllerRequests.PartnerCreateReadingOrderRequest
	if err := commonUtils.RequestBodyParser(c, &request); err != nil {
		return _error.HandleSystemError(c, _error.New(err))
	}

	teleradPartnerUuid := secure.GetUserUuidFromJwt(c)

	result, systemErr := services.PartnerCreateReadingOrder(c.Context(), teleradPartnerUuid, request)
	if systemErr != nil {
		return _error.HandleSystemError(c, systemErr)
	}

	return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "success", result))
}
