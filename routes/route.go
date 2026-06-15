package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	selfCheckRouter "telerad-core-module/routes/health-check"

	"telerad-core-module/internals/routes"
)

var BASE_PATH string = "/services/telerad-core"
var PVT_PATH string = "-internal"

func SetupRoutes(app *fiber.App) {
	app.Get(BASE_PATH, func(c *fiber.Ctx) error {
		claimData := c.Locals("jwtClaims")
		if claimData == nil {
			return c.SendString("Jwt was bypassed")
		}
		return c.JSON(claimData)
	})

	// Middleware đặt locals:logBody chỉ khi là JSON
	setLogBodyIfJSON := func(c *fiber.Ctx) error {
		// Fiber hỗ trợ c.Is("json") để kiểm tra Content-Type application/json (kể cả có charset)
		if c.Is("json") {
			c.Locals("logBody", string(c.Body()))
		} else {
			// Không log body cho multipart/form-data, octet-stream, v.v.
			c.Locals("logBody", "")
		}
		return c.Next()
	}

	// Group api calls with param '/tintuc/api'
	//api := app.Group(BASE_PATH, logger.New())
	loggerConfig := logger.ConfigDefault
	loggerConfig.Format = "${time} | ${status} | ${latency} | ${method} | ${path} | requestBody: ${locals:logBody} | responseBody: ${resBody}\n"
	api := app.Group(BASE_PATH, setLogBodyIfJSON, logger.New(loggerConfig))
	apiPvt := app.Group(BASE_PATH+PVT_PATH, setLogBodyIfJSON, logger.New(loggerConfig))

	// Setup note routes, can use same syntax to add routes for more models
	selfCheckRouter.SetupRoutes(app)
	routes.SetupRoutes(api, apiPvt)

	// catch 404 and forward to error handler
	// app.use(function(req, res, next) {
	// 	next(createError(404));
	// });

}
