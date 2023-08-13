package resp

import (
	"github.com/gofiber/fiber/v2"
)

func Success(ctx *fiber.Ctx) error {
	return success(ctx, struct{}{}, 0, false)
}
func SuccessWithData(ctx *fiber.Ctx, data any) error {
	return success(ctx, data, 0, false)
}
func SuccessArray(ctx *fiber.Ctx, total int, data any) error {
	return success(ctx, data, total, true)
}

func DoNothing(ctx *fiber.Ctx, msg string, isArray ...bool) error {
	return customResponse(ctx, fiber.StatusAccepted, CodeSuccess, nil, msg, isArray)
}
