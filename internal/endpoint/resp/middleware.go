package resp

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorPanic(ctx *fiber.Ctx) error {
	return customResponse(ctx, fiber.StatusInternalServerError, CodeServerPanic, nil, "server panic", nil)
}

func ErrorHost(ctx *fiber.Ctx) error {
	return customResponse(ctx, fiber.StatusForbidden, CodeNoSuchHost, nil, "no such host", nil)
}
