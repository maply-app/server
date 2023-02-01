package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"maply/core/validators"
	"maply/models"
)

func RegisterSerializer(c *fiber.Ctx) (*models.User, bool) {
	data := &models.User{}
	if err := c.BodyParser(data); err != nil {
		return data, false
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return data, false
	}

	// Validate a password
	if !validators.Password(data.Password) {
		return data, false
	}
	return data, true
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=8,max=24"`
}

func LoginSerializer(c *fiber.Ctx) (*LoginInput, bool) {
	data := &LoginInput{}
	if err := c.BodyParser(data); err != nil {
		return data, false
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return data, false
	}
	return data, true
}
