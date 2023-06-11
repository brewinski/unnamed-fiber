package handler

import (
	velocity_api "github.com/brewinski/unnamed-fiber/api/velocity-api"
	"github.com/gofiber/fiber/v2"
)

func PointsEarnHandler(c *fiber.Ctx) error {
	bodyParams := velocity_api.VelocityEarPointsParams{}
	err := c.BodyParser(&bodyParams)

	if err != nil {
		return fiber.ErrBadRequest
	}
	res, err := velocity_api.AllocatePoint(bodyParams)

	if err != nil {
		return err
	}
	return c.JSON(res)
}
