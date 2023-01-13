package core

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status     string    `json:"status"`
	StatusCode int       `json:"-"`
	Info       ErrorInfo `json:"info"`
}

func Error(err ErrorResponse) Response {
	return CreateResponse(err.StatusCode, err)
}

type SuccessResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func Success(statusCode int, data any) Response {
	success := SuccessResponse{}
	success.Status = "success"
	success.Data = data
	return CreateResponse(statusCode, success)
}

type Response struct {
	code int
	data any
}

func CreateResponse(code int, data any) Response {
	response := Response{}
	response.code = code
	response.data = data
	return response
}

func Send(c *fiber.Ctx, response Response) error {
	return c.Status(response.code).JSON(response.data)
}
