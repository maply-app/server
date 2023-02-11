package auth

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	authSerializers "maply/api/serializers/auth"
	"maply/services/auth"
	"net/http"
)

func Register(c *fiber.Ctx) error {
	input, status := authSerializers.RegisterSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	userId, err := auth.CreateUser(input)
	if err != nil {
		return core.Send(c, core.Error(core.ValidationError))
	}
	return core.Send(c, core.Success(http.StatusOK, fiber.Map{"id": userId}))
}

func Login(c *fiber.Ctx) error {
	input, status := authSerializers.LoginSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	t, err := auth.GenerateToken(input.Email, input.Password)

	if err != nil {
		return core.Send(c, core.Error(core.Unauthorized))
	}
	return core.Send(c, core.Success(http.StatusOK, fiber.Map{"token": t}))
}
