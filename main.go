package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"telerad-core-module/configs"
	_ "telerad-core-module/docs"
	"telerad-core-module/internals/repositories"
	"telerad-core-module/internals/responses"
	"telerad-core-module/internals/services"
	"telerad-core-module/jwtchecker"
	routes "telerad-core-module/routes"
	"time"

	"github.com/BeeTechHub/go-common/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

var _currentVersion int64

// @title           Swagger Monitor API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /monitor/api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	logger.InitLogger(configs.GetLogLevel(), configs.GetWriteLogToFile())
	logger.Infof("Starting on time:%s", time.Now().Format("2006-January-02"))

	// https://github.com/gofiber/fiber/issues/623
	// app := fiber.New(&fiber.Settings{
	// 	StrictRouting: true,
	// })
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendString(fmt.Sprintf("Http Code: %d", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})

	// app.Use(cors.New())
	// Configure CORS
	allowedOriginsStr := configs.GetCorsAllowedOrigins()
	corsConfig := cors.Config{
		AllowMethods:  configs.GetCorsAllowMethods(),
		AllowHeaders:  configs.GetCorsAllowHeaders(),
		ExposeHeaders: configs.GetCorsExposeHeaders(),
	}

	if allowedOriginsStr == "" {
		// Default to wildcard '*' if not configured (credentials must be false with wildcard)
		corsConfig.AllowOrigins = "*"
		corsConfig.AllowCredentials = false
	} else {
		// Parse comma-separated origins into slice
		allowedOriginsList := strings.Split(allowedOriginsStr, ",")
		for i := range allowedOriginsList {
			allowedOriginsList[i] = strings.TrimSpace(allowedOriginsList[i])
		}
		// Create a map for faster lookup
		allowedOriginsMap := make(map[string]bool)
		for _, origin := range allowedOriginsList {
			if origin != "" {
				allowedOriginsMap[origin] = true
			}
		}
		// When credentials are included, Access-Control-Allow-Origin cannot be '*'
		corsConfig.AllowCredentials = true
		corsConfig.AllowOriginsFunc = func(origin string) bool {
			// Check if origin is in allowed list
			return allowedOriginsMap[origin]
		}
	}
	app.Use(cors.New(corsConfig))

	jwtchecker.InitPrivateKeyFromContent(configs.GetJwtPrivateKey())
	jwtchecker.InitPublicKeyFromContent(configs.GetJwtPublicKey())

	logger.Infof("connect db start")
	//run database
	configs.ConnectDB()
	repositories.InitDatabase()
	services.InitNoTransaction()

	//init redis
	configs.ConnectRedis()

	logger.Infof("setup route start")
	// Setup the router
	routes.SetupRoutes(app)

	app.Get("/services/category/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Status(http.StatusOK).JSON(responses.NewBaseResponse(http.StatusOK, "Telerad Core service DEMO", time.Now().String()))
	})

	productPort := configs.GetProductPort()
	err := app.Listen(":" + productPort)
	if err != nil {
		logger.Fatalf("fiber.Listen failed %s", err)
	}
}

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}
