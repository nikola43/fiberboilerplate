package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/controllers"
)

func AuthRoutes(router fiber.Router) {
	// /api/v1/auth
	rg := router.Group("/auth")

	// /api/v1/auth/login
	rg.Post("/login", controllers.Login)

	// /api/v1/auth/signup
	rg.Post("/signup", controllers.Signup)
}
