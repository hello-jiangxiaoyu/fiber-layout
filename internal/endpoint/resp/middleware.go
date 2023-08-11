package resp

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ErrorPanic(ctx *fiber.Ctx) error {
	return errorResponse(ctx, http.StatusInternalServerError, ServerPanic, nil, "server panic", nil)
}

func ErrorHost(ctx *fiber.Ctx) error {
	return errorResponse(ctx, http.StatusForbidden, CodeNoSuchHost, nil, "no such host", nil)
}
