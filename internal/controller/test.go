package controller

import (
	"errors"
	"fiber/internal/endpoint/resp"
	"github.com/gofiber/fiber/v2"
)

// SayHello
// @Summary	say hello
// @Schemes
// @Description	返回hello world，添加自定义响应头
// @Tags		fiber-layout
// @Success		200
// @Router		/hello [get]
func SayHello(c *fiber.Ctx) error {
	c.Append("Append-Header", "Fiber-Layout")
	return resp.SuccessWithData(c, fiber.Map{"hello": "fiber"})
}

// GetPostBody
// @Summary	test post body
// @Schemes
// @Description	直接返回请求body内容
// @Tags		fiber-layout
// @Success		200
// @Router		/body [post]
func GetPostBody(c *fiber.Ctx) error {
	return c.Send(c.Body())
}

// GetError
// @Summary	test error
// @Schemes
// @Description	handler返回错误，打印access日志
// @Tags		fiber-layout
// @Success		200
// @Router		/error [get]
func GetError(c *fiber.Ctx) error {
	if err := c.Status(fiber.StatusOK).SendString("error"); err != nil {
		return err
	}
	return errors.New("my error response")
}

// WrapErrorLog
// @Summary	path arg
// @Schemes
// @Description	打印wrap error测试
// @Tags		fiber-layout
// @Success		200
// @Router		/error/wrap [get]
func WrapErrorLog(c *fiber.Ctx) error {
	return resp.ErrorUnknown(c, errors.New("new error"), "my msg", true)
}

// GetPanic
// @Summary	test panic
// @Schemes
// @Description	空指针panic，测试是否能recover，并返回响应
// @Tags		fiber-layout
// @Success		200
// @Router		/panic [get]
func GetPanic(c *fiber.Ctx) error {
	var pt *int
	*pt = 1
	return c.Status(fiber.StatusOK).SendString("ok")
}

// GetArg
// @Summary	path arg
// @Schemes
// @Description	返回路劲参数
// @Tags		fiber-layout
// @Param		appId		path	string	true	"app id"
// @Success		200
// @Router		/args/{arg} [get]
func GetArg(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString(c.Params("arg"))
}
