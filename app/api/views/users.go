package views

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	"maply/api/serializers"
	"maply/services/users"
	"maply/services/validators"
	"net/http"
)

// Get ...
func Get(c *fiber.Ctx) error {
	user, err := users.GetUser(c.Locals("user").(string))

	if err != nil {
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
	return core.Send(c, core.Success(http.StatusOK, user))
}

// Find ...
func Find(c *fiber.Ctx) error {
	input, status := serializers.FindUserSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	users, err := users.FindUser(input.Username)
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, users))
}

// GetByID ...
func GetByID(c *fiber.Ctx) error {
	input := c.Query("userID", "")
	if !validators.UUID(input) {
		return core.Send(c, core.Error(core.ValidationError))
	}

	user, err := users.GetUserByID(input)
	if err != nil {
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
	return core.Send(c, core.Success(http.StatusOK, user))
}
