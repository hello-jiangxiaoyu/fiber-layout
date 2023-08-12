package resp

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// ErrorUnknown 未知错误
func ErrorUnknown(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	return customResponse(ctx, http.StatusInternalServerError, CodeUnknown, err, respMsg, isArray)
}

// ErrorSqlModify SQL修改失败
func ErrorSqlModify(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	return customResponse(ctx, http.StatusInternalServerError, CodeSqlModify, err, respMsg, isArray)
}

// ErrorSelect 数据库查询错误
func ErrorSelect(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	return customResponse(ctx, http.StatusInternalServerError, CodeSqlSelect, err, respMsg, isArray)
}

// ErrorNotFound 资源未找到
func ErrorNotFound(ctx *fiber.Ctx, err error, respMsg string, isArray ...bool) error {
	return customResponse(ctx, http.StatusInternalServerError, CodeNotFound, err, respMsg, isArray)
}

func ErrorSaveSession(ctx *fiber.Ctx, err error, isArray ...bool) error {
	return customResponse(ctx, http.StatusInternalServerError, CodeSaveSession, err, "failed to save session", isArray)
}
