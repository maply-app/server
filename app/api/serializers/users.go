package serializers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FindUserInput struct {
	Username string `validate:"required,min=1,max=24"`
}

func FindUserSerializer(c *fiber.Ctx) (*FindUserInput, bool) {
	username := c.Query("username", "")
	data := &FindUserInput{
		username,
	}
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return data, false
	}
	return data, true
}
