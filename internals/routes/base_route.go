package routes

import "github.com/gofiber/fiber/v2"

var v1 fiber.Router
var v2 fiber.Router
var pvt fiber.Router
var forInternal fiber.Router

const user = "/user"
const admin = "/admin"
const staff = "/staff"
const patient = "/patient"
const public = "/public"
const localGatewayServer = "/lgs"
const integration = "/integration"
const partner = "/partner"

func SetupRoutes(baseRoute fiber.Router, basePvtRoute fiber.Router) {
	v1 = baseRoute.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	v2 = baseRoute.Group("/v2", func(c *fiber.Ctx) error {
		c.Set("Version", "v2")
		return c.Next()
	})

	forInternal = baseRoute.Group("/for-internal", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	pvt = basePvtRoute

	// Đăng ký route theo domain tại đây khi thêm dần.
	AuthRoutes()
	TeleradPartnerRoutes()
	StaffAccountRoutes()
	TeleradReadingOrderRoutes()
	ImagingResultTemplateRoutes()
	ImagingResultSheetTemplateRoutes()
}

func getRoute(route fiber.Router, collection string) fiber.Router {
	return route.Group(collection)
}
