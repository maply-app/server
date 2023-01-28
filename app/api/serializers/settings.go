package serializers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"maply/models"
)

func SettingsSerializer(c *fiber.Ctx) (*models.Settings, bool) {
	data := &models.Settings{}
	if err := c.BodyParser(data); err != nil {
		fmt.Printf("err –> %s", err)
		return data, false
	}

	avatar, err := c.FormFile("avatar")
	if err == nil {
		data.Avatar = avatar
	}

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return data, false
	}
	return data, true
}
