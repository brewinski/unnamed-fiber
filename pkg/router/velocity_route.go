package router

import (
	"github.com/brewinski/unnamed-fiber/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func setupVelocityRoutes(router fiber.Router) {
	velocity := router.Group("velocity")
	velocity.Post("/points-earn", handler.PointsEarnHandler)
}
