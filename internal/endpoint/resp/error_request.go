package resp

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// ErrorRequest 请求参数错误
func ErrorRequest(ctx *fiber.Ctx, err error, msg string, isArray ...bool) error {
	return customResponse(ctx, http.StatusBadRequest, CodeRequestPara, err, msg, isArray)
}

// ErrorForbidden 无权访问
func ErrorForbidden(ctx *fiber.Ctx, msg string, isArray ...bool) error {
	return customResponse(ctx, http.StatusForbidden, CodeForbidden, nil, msg, isArray)
}

// ErrorNoLogin 用户未登录
func ErrorNoLogin(ctx *fiber.Ctx, isArray ...bool) error {
	return customResponse(ctx, http.StatusUnauthorized, CodeNotLogin, nil, "user not login", isArray)
}
