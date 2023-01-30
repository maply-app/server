package users

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	usersSerializers "maply/api/serializers/users"
	"maply/services/users"
	"maply/services/validators"
	"net/http"
)

func Get(c *fiber.Ctx) error {
	u, err := users.GetUser(c.Locals("user").(string))

	if err != nil {
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
	return core.Send(c, core.Success(http.StatusOK, u))
}

func Find(c *fiber.Ctx) error {
	input, status := usersSerializers.FindUserSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	u, err := users.FindUser(input.Username)
	if err != nil {
		return core.Send(c, core.Error(core.InternalServerError))
	}
	return core.Send(c, core.Success(http.StatusOK, u))
}

func GetByID(c *fiber.Ctx) error {
	input := c.Query("userId", "")
	if !validators.UUID(input) {
		return core.Send(c, core.Error(core.ValidationError))
	}

	u, err := users.GetUserByID(input)
	if err != nil {
		return core.Send(c, core.Error(core.ObjectNotFound))
	}
	return core.Send(c, core.Success(http.StatusOK, u))
}
