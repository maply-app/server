package views

import (
	"github.com/gofiber/fiber/v2"
	"maply/api/core"
	"maply/api/serializers"
	"maply/services/auth"
	"net/http"
)

func Register(c *fiber.Ctx) error {
	// Validate request data
	input, status := serializers.RegisterSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	// Insert a user into the database
	id, err := auth.CreateUser(input)
	if err != nil {
		return core.Send(c, core.Error(core.ValidationError))
	}
	return core.Send(c, core.Success(http.StatusOK, fiber.Map{"id": id}))
}

func Login(c *fiber.Ctx) error {
	// Validate request data
	input, status := serializers.LoginSerializer(c)
	if !status {
		return core.Send(c, core.Error(core.ValidationError))
	}

	// Insert a user into the database
	token, err := auth.GenerateToken(input.Email, input.Password)
	if err != nil {
		return core.Send(c, core.Error(core.Unauthorized))
	}
	return core.Send(c, core.Success(http.StatusOK, fiber.Map{"token": token}))
}
