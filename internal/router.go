package internal

import (
	"fiber/internal/controller"
	"fiber/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func NewRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true, // CPU亲和性，Goland断点调试必须设为false !!!!!!
		CaseSensitive: true,
		StrictRouting: true, // 例如"/foo"和"/foo/" 认为是不同的路由
		ServerHeader:  "Fiber-Layout",
		AppName:       "v0.0.1",
	})

	app.Use(pprof.New(pprof.Config{Prefix: "/v1"})) // "/v1/debug/pprof/"

	app.Use(middleware.Recovery(), middleware.GenerateRequestID(), middleware.RequestLogger())
	app.Get("/hello", controller.SayHello)
	app.Get("/error", controller.GetError)
	app.Post("/body", controller.GetPostBody)
	app.Get("/panic", controller.GetPanic)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"msg": "no such route"})
	})

	return app
}
