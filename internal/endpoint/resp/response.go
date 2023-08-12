package resp

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"net/http"
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

func customResponse(c *fiber.Ctx, code int, errCode uint, err error, msg string, isArray []bool) error {
	isArrayFlag := len(isArray) != 0
	var responseData any
	if isArrayFlag {
		responseData = struct{}{}
	} else {
		responseData = []struct{}{}
	}

	if err1 := response(c, code, errCode, msg, responseData, 0, isArrayFlag); err1 != nil {
		if err == nil {
			return err1
		}
		err = errors.WithMessage(err, err1.Error())
	}

	return err
}

func success(c *fiber.Ctx, data any, total int, isArray bool) error {
	return response(c, http.StatusOK, CodeSuccess, MsgSuccess, data, total, isArray)
}
