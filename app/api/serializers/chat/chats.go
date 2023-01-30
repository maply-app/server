package chat

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type GetChatsInput struct {
	Count  int `validate:"min=0,max=25"`
	Offset int `validate:"min=0"`
}

func GetChatsSerializer(c *fiber.Ctx) (GetChatsInput, bool) {
	count, err := strconv.Atoi(c.Query("count", "50"))
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
