package internal

import "github.com/gofiber/fiber/v2"

type Api struct {
	c     *fiber.Ctx
	Error error
}

func (a *Api) SetCtx(ctx *fiber.Ctx) *Api {
	if a.c == nil {
		a.c = ctx
	}
	return a
}

func (a *Api) setError(err error) *Api {
	if a.Error == nil {
		a.Error = err
	}
	return a
}
