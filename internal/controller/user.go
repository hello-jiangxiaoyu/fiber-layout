package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func SayHello(c *fiber.Ctx) error {
	c.Append("Append-Header", "Fiber-layout")
	return c.SendString("Hello world!")
}

func GetPostBody(c *fiber.Ctx) error {
	return c.Send(c.Body())
}

func GetErrorResponse(c *fiber.Ctx) error {
	if err := c.Status(fiber.StatusOK).SendString("error"); err != nil {
		return err
	}
	return errors.New("my error response")
}

func GetPanic(c *fiber.Ctx) error {
	var pt *int
	*pt = 1
	return c.Status(fiber.StatusOK).SendString("ok")
}
