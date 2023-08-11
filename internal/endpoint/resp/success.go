package resp

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const (
	CodeSuccess = 0
	CodeAccept  = 0
	MsgSuccess  = ""
)

type Response struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type ArrayResponse struct {
	Code  uint   `json:"code"`
	Msg   string `json:"msg"`
	Total int    `json:"total"`
	Data  any    `json:"data"`
}

func response(c *fiber.Ctx, code int, errCode uint, msg string, data any, total int, isArray bool) error {
	if isArray {
		return c.Status(code).JSON(&ArrayResponse{Code: errCode, Msg: msg, Total: total, Data: data})
	}

	return c.Status(code).JSON(&Response{Code: errCode, Msg: msg, Data: data})
}

func success(c *fiber.Ctx, data any, total int, isArray bool) error {
	return response(c, http.StatusOK, CodeSuccess, MsgSuccess, data, total, isArray)
}

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
	return response(ctx, http.StatusAccepted, CodeSuccess, msg, nil, 0, len(isArray) != 0)
}
