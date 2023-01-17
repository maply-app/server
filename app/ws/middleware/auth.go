package middleware

import (
	"github.com/gofiber/fiber/v2"
	"maply/errors"
	"maply/services/utils"
)

func UserIdentity(c *fiber.Ctx) error {
	token := c.Query("Token")
	if token == "" {
		return errors.Unauthorized
	}

	t, err := utils.ParseToken(token)
	if err != nil {
		return errors.Unauthorized
	}

	c.Locals("userId", t)
	return nil
}
