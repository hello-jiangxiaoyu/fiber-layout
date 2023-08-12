package middleware

import (
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"strconv"
)

func GenerateRequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID, ok := c.GetReqHeaders()[fiber.HeaderXRequestID]
		if !ok {
			requestID = strconv.Itoa(rand.Int())
			c.Set(fiber.HeaderXRequestID, requestID)
		}
		return c.Next()
	}
}
