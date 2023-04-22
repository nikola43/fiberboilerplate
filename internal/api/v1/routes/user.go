package routes

import (
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/auth"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/controllers"
)

func UserRoutes(router fiber.Router) {
	// /api/v1/user
	rg := router.Group("/user")

	// JWT Middleware
	rg.Use(jwtware.New(jwtware.Config{
		SigningMethod: "RS256",
		SigningKey:    auth.PrivateKey.Public(),
	}))

	// /api/v1/user/{id}
	rg.Get("/:id", controllers.GetUserById)

	// /api/v1/user/{id}
	rg.Put("/:id", controllers.Update)

	// /api/v1/user/{id}
	rg.Delete("/:id", controllers.Delete)
}
