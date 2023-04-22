package app

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	//"github.com/gofiber/helmet"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/auth"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/config"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/database"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/routes"
	"github.com/nikola43/fiberboilerplate/pkg/utils"
)

var httpServer *fiber.App

func Initialize() {
	apiConfig, err := config.ReadConfig("config.yaml")
	if err != nil {
		log.Fatal("Error reading config file, %s", err)
	}

	privateKey, err := utils.ReadRsaPrivateKeyFromFile()
	if err != nil {
		log.Fatal("Error reading private key, %s", err)
	}
	auth.PrivateKey = privateKey

	// initialize database
	database.InitializeDatabase(apiConfig.Database)

	// migrate database
	//database.Migrate()

	// initialize http server
	httpServer = fiber.New(fiber.Config{
		BodyLimit: 2000 * 1024 * 1024, // this is the default limit of 2GB
	})
	/*
		//httpServer.Use(middlewares.XApiKeyMiddleware)
		httpServer.Use(cors.New(cors.Config{
			AllowOrigins: "https://web.com",
		}))
	*/
	httpServer.Use(cors.New(cors.Config{}))
	httpServer.Use(logger.New())
	//httpServer.Use(helmet.New())

	api := httpServer.Group("/api") // /api
	v1 := api.Group("/v1")          // /api/v1
	handleRoutes(v1)

	// print server url
	fmt.Println(color.YellowString("  ----------------- Server Info -----------------"))
	fmt.Println(color.CyanString("http://localhost:" + apiConfig.Port))

	err = httpServer.Listen(":" + apiConfig.Port)
	if err != nil {
		log.Fatal("Error starting server, %s", err)
	}
}

func handleRoutes(api fiber.Router) {
	routes.AuthRoutes(api)
	routes.UserRoutes(api)
	//routes.InverterRoutes(api)
}
