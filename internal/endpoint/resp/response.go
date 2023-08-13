package resp

import (
	"fiber/pkg/utils"
	"github.com/gofiber/fiber/v2"
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
	c.Locals("code", errCode)
	if isArray {
		return c.Status(code).JSON(&ArrayResponse{Code: errCode, Msg: msg, Total: total, Data: data})
	}

	return c.Status(code).JSON(&Response{Code: errCode, Msg: msg, Data: data})
}

func customResponse(c *fiber.Ctx, code int, errCode uint, err error, msg string, isArray []bool) error {
	isArrayFlag := len(isArray) != 0
	var responseData any
	if isArrayFlag {
		responseData = []struct{}{}
	} else {
		responseData = struct{}{}
	}

	err = utils.WithStringError(err, msg)
	err = utils.WrapError(err, response(c, code, errCode, msg, responseData, 0, isArrayFlag))

	return err
}

func success(c *fiber.Ctx, data any, total int, isArray bool) error {
	return response(c, fiber.StatusOK, CodeSuccess, MsgSuccess, data, total, isArray)
}
