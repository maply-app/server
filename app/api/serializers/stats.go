package serializers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"maply/models"
)

func StatsSerializer(c *fiber.Ctx) (*models.Stats, bool) {
	data := &models.Stats{}
	if err := c.BodyParser(data); err != nil {
		return data, false
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return data, false
	}
	return data, true
}
