package serializers

import (
	"github.com/gofiber/fiber/v2"
	"maply/models"
)

func SendRequestSerializer(c *fiber.Ctx) (*models.Request, bool) {
	data := &models.Request{}
	if err := c.BodyParser(data); err != nil {
		return data, false
	}

	data.SenderID = c.Locals("user").(string)
	if data.SenderID == data.ReceiverID {
		return data, false
	}
	return data, true
}
