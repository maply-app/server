package serializers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"maply/models"
	"strconv"
)

func SendMessageSerializer(c *fiber.Ctx) (*models.Message, bool) {
	data := &models.Message{}
	if err := c.BodyParser(data); err != nil {
		return data, false
	}

	data.SenderID = c.Locals("user").(string)
	if data.SenderID == data.ReceiverID {
		return data, false
	}
	return data, true
}

type GetMessagesInput struct {
	ReceiverId string `validate:"required,uuid"`
	Count      int    `validate:"min=0,max=25"`
	Offset     int    `validate:"min=0"`
}

func GetMessagesSerializer(c *fiber.Ctx) (GetMessagesInput, bool) {
	receiverId := c.Query("receiverId", "")
	count, err := strconv.Atoi(c.Query("count", "25"))
	if err != nil {
		return GetMessagesInput{}, false
	}
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return GetMessagesInput{}, false
	}
	data := GetMessagesInput{
		receiverId,
		count,
		offset,
	}

	validate := validator.New()
	if err = validate.Struct(data); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
		return data, false
	}
	return data, true
}

type GetChatsInput struct {
	Count  int `validate:"min=0,max=25"`
	Offset int `validate:"min=0"`
}

func GetChatsSerializer(c *fiber.Ctx) (GetChatsInput, bool) {
	count, err := strconv.Atoi(c.Query("count", "25"))
	if err != nil {
		return GetChatsInput{}, false
	}
	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return GetChatsInput{}, false
	}
	data := GetChatsInput{
		count,
		offset,
	}

	validate := validator.New()
	if err = validate.Struct(data); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
		return data, false
	}
	return data, true
}
