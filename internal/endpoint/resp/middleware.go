package resp

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ErrorPanic(ctx *fiber.Ctx) error {
	return customResponse(ctx, http.StatusInternalServerError, CodeServerPanic, nil, "server panic", nil)
}

func ErrorHost(ctx *fiber.Ctx) error {
	return customResponse(ctx, http.StatusForbidden, CodeNoSuchHost, nil, "no such host", nil)
}
