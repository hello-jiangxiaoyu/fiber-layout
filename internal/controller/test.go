package controller

import (
	"errors"
	"fiber/internal/controller/internal"
	"fiber/internal/endpoint/resp"
	"fiber/pkg/client"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type TestRoute struct {
	internal.Api
}

func NewTestRoute() *TestRoute {
	return &TestRoute{}
}

// SayHello
// @Summary	say hello
// @Schemes
// @Description	返回hello world，添加自定义响应头
// @Tags		fiber-layout
// @Success		200
// @Router		/ [get]
func (t TestRoute) SayHello(c *fiber.Ctx) error {
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
func (t TestRoute) GetPostBody(c *fiber.Ctx) error {
	return c.Send(c.Body())
}

// GetError
// @Summary	test error
// @Schemes
// @Description	handler返回错误，打印access日志
// @Tags		fiber-layout
// @Success		200
// @Router		/error [get]
func (t TestRoute) GetError(c *fiber.Ctx) error {
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
func (t TestRoute) WrapErrorLog(c *fiber.Ctx) error {
	return resp.ErrorUnknown(c, errors.New("new error"), "my msg", true)
}

// GetPanic
// @Summary	test panic
// @Schemes
// @Description	空指针panic，测试是否能recover，并返回响应
// @Tags		fiber-layout
// @Success		200
// @Router		/panic [get]
func (t TestRoute) GetPanic(c *fiber.Ctx) error {
	var pt *int
	*pt = 1
	return c.Status(fiber.StatusOK).SendString("ok")
}

// GetArg
// @Summary	path arg
// @Schemes
// @Description	返回路劲参数
// @Tags		fiber-layout
// @Param		arg		path	string	true	"arg"
// @Success		200
// @Router		/args/{arg} [get]
func (t TestRoute) GetArg(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString(c.Params("arg"))
}

// BindJson
// @Summary	bind json test
// @Schemes
// @Description	绑定uri和body参数，并进行参数校验
// @Tags		fiber-layout
// @Param		arg		path	string	true	"arg"
// @Success		200
// @Router		/bind/{arg} [post]
func (t TestRoute) BindJson(c *fiber.Ctx) error {
	type Person struct {
		Arg  string `json:"-" params:"arg" validate:"required"`
		Age  int    `json:"age" params:"-"`
		Name string `json:"name" params:"-" validate:"required"`
	}
	var p Person
	if err := t.SetCtx(c).BindUriAndJson(&p).Error; err != nil {
		return resp.ErrorRequest(c, err, "bind test err")
	}

	return resp.SuccessWithData(c, &p)
}

// GetFastHttpRequest
// @Summary	send fast http request
// @Schemes
// @Description	发送fast http请求
// @Tags		fiber-layout
// @Success		200
// @Router		/fast [get]
func (t TestRoute) GetFastHttpRequest(c *fiber.Ctx) error {
	var res fiber.Map
	if err := client.Get("http://localhost/api/quick/apps", &res); err != nil {
		return resp.ErrorSendRequest(c, err, "get fast err")
	}
	return resp.SuccessWithData(c, &res)
}

// PostFastHttpRequest
// @Summary	send fast http get request
// @Schemes
// @Description	发送fast http get请求
// @Tags		fiber-layout
// @Success		200
// @Router		/fast [post]
func (t TestRoute) PostFastHttpRequest(c *fiber.Ctx) error {
	var res fiber.Map
	data := map[string]any{
		"post": "data",
	}
	if err := client.Post("http://localhost:8000/body", &data, &res); err != nil {
		return resp.ErrorSendRequest(c, err, "post fast err")
	}
	return resp.SuccessWithData(c, &res)
}
