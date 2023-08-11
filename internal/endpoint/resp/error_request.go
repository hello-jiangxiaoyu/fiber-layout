package resp

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const (
	CodeRequestPara = 1000
	CodeForbidden   = 1001
	CodeNoLogin     = 1002
	CodeNoSuchHost  = 1003
)

func errorResponse(c *fiber.Ctx, code int, errCode uint, err error, msg string, isArray []bool) error {

	if len(isArray) == 0 {
		return response(c, code, errCode, msg, struct{}{}, 0, false)
	} else {
		return response(c, code, errCode, msg, struct{}{}, 0, true)
	}

}

// ErrorRequest 请求参数错误
func ErrorRequest(ctx *fiber.Ctx, err error, msg string, isArray ...bool) error {
	return errorResponse(ctx, http.StatusBadRequest, CodeRequestPara, err, msg, isArray)
}

// ErrorForbidden 无权访问
func ErrorForbidden(ctx *fiber.Ctx, msg string, isArray ...bool) error {
	return errorResponse(ctx, http.StatusForbidden, CodeForbidden, nil, msg, isArray)
}

// ErrorNoLogin 用户未登录
func ErrorNoLogin(ctx *fiber.Ctx, isArray ...bool) error {
	return errorResponse(ctx, http.StatusUnauthorized, CodeNoLogin, nil, "user not login", isArray)
}
