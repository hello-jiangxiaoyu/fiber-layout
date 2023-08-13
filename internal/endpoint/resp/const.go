package resp

import (
	"github.com/gofiber/fiber/v2"
)

const (
	MsgSuccess = ""
)

const (
	CodeSuccess = fiber.StatusOK
	CodeAccept  = fiber.StatusAccepted

	CodeNoSuchRoute = 1000 // 系统相关错误码
	CodeRequestPara = 1001
	CodeForbidden   = 1002
	CodeNoRoute     = 1003

	CodeNoSuchHost   = 2000 // 业务相关错误码
	CodeNotLogin     = 2001
	CodeInvalidToken = 2002

	CodeSqlSelect          = 3000 // sql相关错误码
	CodeSqlModify          = 3001
	CodeSqlCreate          = 3002
	CodeSqlDelete          = 3003
	CodeSqlCreateDuplicate = 3004

	CodeServerPanic = 5000 // error相关错误码
	CodeUnknown     = 5001
	CodeNotFound    = 5002
	CodeSaveSession = 5003
	CodeSendRequest = 5004
)
