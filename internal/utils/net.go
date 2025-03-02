package utils

import "github.com/gofiber/fiber/v2"

func Serialize[T any](c *fiber.Ctx) (T, error) {
	var req T
	if err := c.BodyParser(&req); err != nil {
		return req, err
	}

	return req, nil
}
