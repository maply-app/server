package users

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"maply/models"
)

func SettingsSerializer(c *fiber.Ctx) (*models.Settings, bool) {
	data := &models.Settings{}
	if err := c.BodyParser(data); err != nil {
		return data, false
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return data, false
	}
	return data, true
}
